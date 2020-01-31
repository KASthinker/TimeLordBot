package localization

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTranslate(t *testing.T) {
	t.Parallel()
	message := Translate("ru_RU", "buttons", "Menu")
	require.Equal(t, message, "Меню")

	message = Translate("ru_RU", "task", "Priority:")
	require.Equal(t, message, "🔥 *Приоритет:* ")

	message = Translate("ru_RU", "message", "Task list:")
	require.Equal(t, message, "*Список задач:*\n")
}