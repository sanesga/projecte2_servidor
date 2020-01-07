import { EfavoriteBookActions } from '../actions/favoriteBook.actions';
import { favoriteBookActions } from '../actions/favoriteBook.actions';
import { initialfavoriteBookState, favoriteBookState } from '../state/favoriteBook.state';

export const favoriteBookReducers = (
  state = initialfavoriteBookState,
  action: favoriteBookActions
): favoriteBookState => {
  switch (action.type) {
    case EfavoriteBookActions.GetfavoriteBooksSuccess: {
      return {
        ...state,
        favoriteBooks: action.payload
      };
    }
    default:
      return state;
  }
};
