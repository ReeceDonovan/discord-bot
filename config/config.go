package config

import (
	"strings"

	"github.com/Strum355/log"
	"github.com/spf13/viper"
)

// Servers represents the servers.
type Servers struct {
	PublicServer    string `json:"public"`
	CommitteeServer string `json:"committee"`
}

// Channels required for events.
type Channels struct {
	PublicAnnouncements string `json:"public_announcements"` // On public server
	PublicGeneral       string `json:"public_general"`       // On public server
	PrivateEvents       string `json:"private_events"`       // On committee server
}

const limitChars = 8

// InitConfig sets up viper and consul.
func InitConfig() error {
	// Viper
	initDefaults()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // For gamers only
	viper.AutomaticEnv()
	viper.Set(
		"discord.servers",
		&Servers{PublicServer: viper.GetString("discord.public.server"), CommitteeServer: viper.GetString("discord.committee.server")},
	)
	viper.Set(
		"discord.channels",
		&Channels{PublicAnnouncements: viper.GetString("discord.public.channel"), PrivateEvents: viper.GetString("discord.committee.channel"), PublicGeneral: viper.GetString("discord.public.general")},
	)
	welcomeMessages := []string{}
	for _, message := range strings.Split(viper.GetString("discord.public.welcome"), ",") {
		welcomeMessages = append(welcomeMessages, message)
	}
	viper.Set("discord.welcome_messages", &welcomeMessages)

	printAll()
	return nil
}

func printAll() {
	store := log.Fields{}
	for k, v := range viper.AllSettings() {
		localStore := map[string]interface{}{}
		for sk, sv := range v.(map[string]interface{}) {
			if strval, ok := sv.(string); ok {
				if len(strval) > limitChars {
					localStore[sk] = strval[:limitChars] + "..."
				} else {
					localStore[sk] = strval
				}
			} else {
				localStore[sk] = sv
			}
		}
		store[k] = localStore
	}
	log.WithFields(store).Info("discord bot startup config values")
}
