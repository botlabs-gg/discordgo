package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"reflect"

	"github.com/jonas747/discordgo"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {

	flag.StringVar(&Token, "t", "", "Account Token")
	flag.Parse()
}

func main() {

	// Create a new Discord session using the provided login information.
	dg, err := discordgo.New(Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	resp, err := dg.GatewayBot()
	fmt.Printf("%#v\n\n%v\n", resp, err)

	// dg.ShardCount = 1536
	// dg.ShardID = 999
	// dg.AddHandler(handleEventMessage)
	// dg.AddHandler(handleEvent)
	// dg.AddHandler(handleReady)
	// dg.AddHandler(handleGuildCreate)
	dg.AddHandler(handleChannelCreate)
	dg.AddHandler(handleChannelUpdate)

	err = dg.Open()
	if err != nil {
		panic(err)
	}

	select {}
}

const CheckingGuild = 614909558585819162

func handleGuildCreate(session *discordgo.Session, r *discordgo.GuildCreate) {
	if r.ID != 609205310132977666 {
		return
	}

	serializedRoles, _ := json.MarshalIndent(r.Roles, "", "  ")
	// serializedChannel, _ := json.MarshalIndent(r.Channels, "", "  ")
	var serializedMember []byte
	var serializedChannel []byte

	for _, v := range r.Members {
		if v.User.ID == 204255221017214977 {
			serializedMember, _ = json.MarshalIndent(v, "", "  ")
		}
	}

	for _, v := range r.Channels {
		if v.ID == 610485390549188644 {
			serializedChannel, _ = json.MarshalIndent(v, "", "  ")
		}
	}

	fmt.Println("Roles: \n", string(serializedRoles))
	fmt.Println("Member: \n", string(serializedMember))
	fmt.Println("Channel: \n", string(serializedChannel))
}

func handleReady(session *discordgo.Session, r *discordgo.Ready) {
	found := false
	for _, v := range r.Guilds {
		if v.ID == CheckingGuild {
			found = true
			break
		}
	}
	fmt.Println("Got ready! guild: ", len(r.Guilds), "found? ", found)
}

func handleEventMessage(session *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.GuildID != CheckingGuild {
		return
	}
	printJson(msg)
	fmt.Println("Got event", msg.GuildID)
}

func handleEvent(session *discordgo.Session, i interface{}) {
	cast, ok := i.(discordgo.GuildEvent)
	if !ok {
		return
	}

	if cast.GetGuildID() != CheckingGuild {
		return
	}
	typ := reflect.TypeOf(cast)

	serialized, _ := json.MarshalIndent(i, "", "  ")
	fmt.Printf("Got event %s\n%s\n", typ.String(), string(serialized))
}

func printJson(in interface{}) {
	serialised, _ := json.MarshalIndent(in, "", "  ")
	fmt.Println(string(serialised))
}

func handleChannelCreate(session *discordgo.Session, r *discordgo.ChannelCreate) {
	fmt.Println("Got channel CREATE: ", r.GuildID, r.Name, r.Position)
}

func handleChannelUpdate(session *discordgo.Session, r *discordgo.ChannelUpdate) {
	fmt.Println("Got channel UPDATE: ", r.GuildID, r.Name, r.Position)
}
