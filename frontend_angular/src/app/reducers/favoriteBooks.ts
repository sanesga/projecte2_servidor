import * as favoriteBooks from '../actions/favoriteBooks';

export function reducer(state: [], action: favoriteBooks.FavoriteBooksUpdatedAction) {
    switch (action.type) {
        case favoriteBooks.FAVORITEBOOKUPDATED:
            return action.payload;        
        default:
            return state;
    }
}