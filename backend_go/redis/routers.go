package redis

import (
	"fmt"
	"net/http"

	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func Routers(router *gin.RouterGroup) {
	router.GET("/", getAll)
	router.GET("/:key", getOne)
	router.POST("/", save)
}

//estructura de datos guardados, clave-valor
type Info struct {
	Key   string `json:"key"   binding:"required"`
	Value string `json:"value" binding:"required"`
}

//obtenemos todos los datos
func getAll(c *gin.Context) {

	client := newClient()
	var array map[string]string
	array = make(map[string]string)

	//si error
	keys, err := client.Do("KEYS", "*").Result()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	//obtenemos los valores de las claves
	for i := 0; i < reflect.ValueOf(keys).Len(); i++ {
		key := fmt.Sprintf("%v", reflect.ValueOf(keys).Index(i)) // convert from interface to string

		err, val := get(key, client)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		array[key] = val
	}
	c.JSON(200, gin.H{"keys": array})
}

//obtenemos un dato, especificando la clave
func getOne(c *gin.Context) {
	client := newClient()

	key := c.Param("key")

	err, val := get(key, client)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{key: val})

}

//guardamos los datos
func save(c *gin.Context) {
	client := newClient()

	var info Info
	if err := c.ShouldBindJSON(&info); err != nil { //marca error, pero funciona
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//pasamos los datos a json
	c.BindJSON(&info)

	//si error
	err := set(info.Key, info.Value, client)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//si se guardan los datos correctamente
	c.JSON(200, gin.H{"result": "ok"})
}

//creamos el cliente redis
func newClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	return client
}

func set(key string, value string, client *redis.Client) error {
	err := client.Set(key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func get(key string, client *redis.Client) (error, string) {
	val, err := client.Get(key).Result()
	return err, val
}
