package core

import (
	"net/http"

	"github.com/discuitnet/discuit/internal/httperr"
)

var (
	// ErrWrongPassword is returned by MatchLoginCredentials if username and password
	// do not match.
	ErrWrongPassword = &httperr.Error{HTTPStatus: http.StatusUnauthorized, Code: "wrong-password", Message: "Имя пользователя и пароль не совпадают."}

	ErrUserDeleted = httperr.NewForbidden("user-deleted", "Не удается продолжить, так как пользователь удален.")
)

var (
	errNotAuthor = httperr.NewForbidden("not_author", "Вы не являетесь автором.")
	errNotMod    = httperr.NewForbidden("not_mod", "Вы не являетесь модератором.")
	errNotAdmin  = httperr.NewForbidden("not_admin", "Вы не являетесь администратором.")

	errImageNotFound = httperr.NewNotFound("image-not-found", "Изображение не найдено.")

	errCommunityNotFound = httperr.NewNotFound("community/not-found", "Сообщество не найдено.")

	errUserNotFound            = httperr.NewNotFound("user_not_found", "Пользователь не найден.")
	errUserBannedFromCommunity = httperr.NewForbidden("banned-from-community", "Пользователь заблокирован в сообществе.")

	errCommentDeleted  = httperr.NewForbidden("comment_deleted", "Комментарий(ы) удален(ы).")
	errCommentNotFound = httperr.NewNotFound("comment_not_found", "Комментарий(ы) не найден.")

	errPostNotFound        = httperr.NewNotFound("post/not-found", "Сообщения не найдены.")
	errPostLocked          = httperr.NewForbidden("post-locked", "Пост закрыт.")
	errPostTypeUnsupported = httperr.NewBadRequest("post-type/unsupported", "Неподдерживаемый тип записи.")

	errInvalidUserGroup = httperr.NewBadRequest("user/invalid-group", "Недопустимая группа пользователей.")
)
