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

	// Set these to the shard you want to debug
	// dg.ShardCount = ...
	// dg.ShardID = ...

	resp, err := dg.GatewayBot()
	fmt.Printf("Gateway bot: %#v\n\n%v\n", resp, err)

	dg.AddHandler(handleReady)
	dg.AddHandler(dumpEvent)

	err = dg.Open()
	if err != nil {
		panic(err)
	}

	fmt.Println("Running, ctrl-c to exit")

	select {}
}

func handleReady(session *discordgo.Session, r *discordgo.Ready) {
	fmt.Println("Got ready! n guilds: ", len(r.Guilds))
}

func dumpEvent(session *discordgo.Session, evt interface{}) {
	switch evt.(type) {
	case *discordgo.GuildCreate /*, *discordgo.PresenceUpdate*/ :
		// this is a pretty heavy event that will spam the terminal for a while so skip this event
		// you can add more cases to skip other spammy events such as presence updates above
		return
	}

	typ := reflect.TypeOf(evt)
	serialized, _ := json.MarshalIndent(evt, "", "  ")
	fmt.Printf("Got event %s\n%s\n", typ.String(), string(serialized))
}
