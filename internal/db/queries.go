package db

import (
	"fmt"
	"strings"

	sqrl "github.com/Masterminds/squirrel"
)

var (
	// используется для генерации sql-запросов к базе данных postgresql
	sb = sqrl.StatementBuilder.PlaceholderFormat(sqrl.Dollar)

	// список полей, возвращаемых с информацией о пользователе
	userFields = []string{"uid", "email", "properties", "updated", "blocked"}
	// используется как подзапрос для получения информации, что email подтвержден
	sbVerifiedField = sqrl.Alias(sb.
			Select("1").
			Prefix("EXISTS (").
			From("emails").
			Where("email = users.email").
			Suffix(")"),
		"verified")
	// заготовка для возврата данных о пользователе в запросе
	sbReturnUser = sqrl.ConcatExpr(
		"RETURNING ", strings.Join(userFields, ", "), ", ", sbVerifiedField)

	// SuffixExpr(sbReturnUser)
	// регистрирует нового пользователя с логином и паролем
	// если такой пользователь уже зарегистрирован, но не имеет пароля, то
	// добавляет ему пароль
	sqlInsertUser = toSQL(sb.
			Insert("users").
			Columns("email", "password").
			Values("", "").
			Suffix("ON CONFLICT (email) DO UPDATE SET password = EXCLUDED.password, updated = DEFAULT, logged = DEFAULT WHERE users.password IS NULL").
			SuffixExpr(sbReturnUser))
	// регистрирует нового пользователя
	// обновляет дату обновления и свойства для уже зарегистрированого
	// пользователя, если они не заданы
	sqlInsertUserOpenID = toSQL(sb.
				Insert("users").
				Columns("email", "properties").
				Values("", nil).
				Suffix("ON CONFLICT (email) DO UPDATE SET properties = COALESCE(users.properties, EXCLUDED.properties), updated = DEFAULT, logged = DEFAULT").
				SuffixExpr(sbReturnUser))

	// заготовка запроса для получения данных о пользователе
	sbSelectUser = sb.
			Select(userFields...).
			Column(sbVerifiedField).
			From("users")
	// заготовка запроса для получения данных о пользователе по email-адресу
	sbSelectUserByEmail = sbSelectUser.
				Where(sqrl.Eq{"email": ""})
	// возвращает информацию о пользователе по email
	sqlSelectUser = toSQL(sbSelectUserByEmail)
	// возвращает информацию о пользователе и его пароль
	sqlSelectPassword = toSQL(sbSelectUserByEmail.
				Column("password").
				Where(sqrl.NotEq{"password": nil}))
	// возвращает информацию о пользователе по openid регистрации
	sqlSelectUserOpenID = toSQL(sbSelectUser.
				Join("openid USING (uid)").
				Where(sqrl.Eq{"provider": ""}).
				Where(sqrl.Eq{"subject": ""}))
	// возвращает информацию о пользователе по email или идентификатору
	sqlSelectUserByEmailOrUID = toSQL(sbSelectUser.
					Where(sqrl.Or{sqrl.Eq{"uid": ""}, sqrl.Eq{"email": ""}}))

	// заготовка для обновления пользователя по его идентификатору
	sbUpdateUser = sb.
			Update("users").
			Set("updated", sqrl.Expr("DEFAULT")).
			Where(sqrl.Eq{"uid": ""})
	// обновляет пароль пользователя
	sqlUpdatePassword = toSQL(sbUpdateUser.
				Set("password", ""))
	// обновляет email и расширенные свойства пользователя
	sqlUpdateUser = toSQL(sbUpdateUser.
			Set("email", "").
			Set("properties", nil))
	// блокирует/разблокирует пользователя
	sqlBlockUser = toSQL(sbUpdateUser.
			Set("blocked", true))
	// обновляет дату последней авторизации пользователя
	sqlLogged = toSQL(sb.
			Update("users").
			Set("logged", sqrl.Expr("DEFAULT")).
			Where(sqrl.Eq{"uid": ""}))
	// сохраняет данные о внешней регистрации пользователя
	sqlInsertOpenID = toSQL(sb.
			Insert("openid").
			Columns("provider", "subject", "uid").
			Values("", "", "").
			Suffix("ON CONFLICT (provider,subject) DO UPDATE SET uid = EXCLUDED.uid"))

	// добавляет email в список подтвержденных
	sqlInsertVerifiedEmail = toSQL(sb.
				Insert("emails").
				Columns("email").
				Values("").
				Suffix("ON CONFLICT (email) DO UPDATE SET updated = DEFAULT"))

	// добавляет в журнал маркетинговую информацию о регистрации
	sqlInsertRegInfo = toSQL(sb.
				Insert("reginfo").
				Columns("domain", "uid", "email", "provider", "referer", "utm").
				Values("", "", "", nil, nil, nil))

	// добавляет новый токен для проверки почты или сброса пароля
	sqlInsertToken = toSQL(sb.
			Insert("tokens").
			Columns("domain", "email", "type").
			Values("", "", 0).
			Suffix("ON CONFLICT (domain, email, type) DO UPDATE SET id=DEFAULT, sended=DEFAULT, created=DEFAULT").
			Suffix("RETURNING id"))
	// удаляет проверочный токен
	sqlDeleteToken = toSQL(sb.
			Delete("tokens").
			Where(sqrl.Eq{"id": ""}).
			Suffix("RETURNING email, type, created"))
	// возвращает список токенов для отправки
	sqlSelectTokens = toSQL(sb.
			Select("id", "domain", "email", "type").
			From("tokens").
			Where("sended = FALSE"))
	sqlUpdateToken = toSQL(sb.
			Update("tokens").
			Set("sended", sqrl.Expr("TRUE")).
			Where(sqrl.Eq{"id": ""}))
)

// toSQL формирует и возвращает строку с sql-запросом.
// Вызывает panic в случае ошибки в запросе.
func toSQL(query sqrl.Sqlizer) string {
	sql, args, err := query.ToSql()
	if err != nil {
		panic(fmt.Errorf("sql query generation error: %w", err))
	}
	_ = args
	// fmt.Printf("%s\nParams: %d\n%s\n", sql, len(args), strings.Repeat("-", 80))
	return sql
}
