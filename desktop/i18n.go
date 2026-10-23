package desktop

import (
	"os"
	"strings"
)

// Lang is a language the interface speaks.
type Lang string

const (
	Portuguese Lang = "pt"
	English    Lang = "en"
)

// DetectLang picks Portuguese when the system locale is Portuguese and
// English otherwise, the same rule the web version applies to the browser.
func DetectLang(getenv func(string) string) Lang {
	for _, key := range []string{"LC_ALL", "LC_MESSAGES", "LANG", "LANGUAGE"} {
		if v := getenv(key); v != "" {
			if strings.HasPrefix(strings.ToLower(v), "pt") {
				return Portuguese
			}
			return English
		}
	}
	return English
}

// SystemLang is DetectLang over the real environment.
func SystemLang() Lang { return DetectLang(os.Getenv) }

// T returns the text for key, or the key itself when it has no translation.
func (l Lang) T(key string) string {
	if s, ok := translations[l][key]; ok {
		return s
	}
	return key
}

var translations = map[Lang]map[string]string{
	Portuguese: {
		"tagline":       "Seu parceiro de foco no trabalho",
		"activeLabel":   "Alerta ativo",
		"noAlert":       "Nenhum alerta selecionado",
		"alertsTitle":   "Meus Alertas",
		"start":         "Iniciar",
		"pause":         "Pausar",
		"reset":         "Resetar",
		"addAlert":      "Adicionar",
		"cancel":        "Limpar",
		"cancelEdit":    "Cancelar",
		"nameLabel":     "Nome do alerta",
		"hoursLabel":    "Horas",
		"minutesLabel":  "Min",
		"secondsLabel":  "Seg",
		"use":           "Usar",
		"edit":          "Editar",
		"save":          "Salvar",
		"delete":        "Excluir",
		"undo":          "Desfazer",
		"ok":            "OK",
		"snooze":        "Adiar 5 min",
		"attention":     "Atenção",
		"invalidAlert":  "Insira um nome válido e um tempo maior que 0.",
		"timeUpMessage": "O tempo do seu alerta finalizou!",
		"hint":          "Dica: aperte Espaço para iniciar ou pausar",
		"added":         "Alerta adicionado",
		"updated":       "Alerta atualizado",
		"deleted":       "Alerta excluído",
		"snoozed":       "Alerta adiado em 5 min",
		"saveFailed":    "Não foi possível salvar os alertas",
		"loadFailed":    "Não foi possível ler os alertas salvos",
		"emptyTitle":    "Nenhum alerta ainda",
		"emptyDesc":     "Crie seu primeiro alerta no formulário acima.",
	},
	English: {
		"tagline":       "Your work focus companion",
		"activeLabel":   "Active alert",
		"noAlert":       "No alert selected",
		"alertsTitle":   "My Alerts",
		"start":         "Start",
		"pause":         "Pause",
		"reset":         "Reset",
		"addAlert":      "Add",
		"cancel":        "Clear",
		"cancelEdit":    "Cancel",
		"nameLabel":     "Alert name",
		"hoursLabel":    "Hours",
		"minutesLabel":  "Min",
		"secondsLabel":  "Sec",
		"use":           "Use",
		"edit":          "Edit",
		"save":          "Save",
		"delete":        "Delete",
		"undo":          "Undo",
		"ok":            "OK",
		"snooze":        "Snooze 5 min",
		"attention":     "Attention",
		"invalidAlert":  "Enter a valid name and a time greater than 0.",
		"timeUpMessage": "Your alert's time has finished!",
		"hint":          "Tip: press Space to start or pause",
		"added":         "Alert added",
		"updated":       "Alert updated",
		"deleted":       "Alert deleted",
		"snoozed":       "Alert snoozed by 5 min",
		"saveFailed":    "Could not save the alerts",
		"loadFailed":    "Could not read the saved alerts",
		"emptyTitle":    "No alerts yet",
		"emptyDesc":     "Create your first alert using the form above.",
	},
}
