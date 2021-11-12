package main
import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

type Games struct {
	Games []User `json:"game"`
}
type Games2 struct {
	Games2 []BJ `json:"game"`
}
type User struct {
	Rps Rps `json:"rps"`
}
type Rps struct {
	Wins int `json:"wins"`
	Losses int `json:Losses`
	Ties int `json:Ties`
}
type BJ struct {
	BlackJack BlackJack `json:"blackjack"`
}
type BlackJack struct {
	Wins int `json:"wins"`
	Losses int `json:Losses`
	Ties int `json:Ties`
}
var playertotal = 0
var enemytotal = 0
var originaluserid = ""
var originalmessageid = ""
func savetoJson(gametype string, wins int, losses int, ties int) {
	fmt.Println(wins, losses, ties)
	switch gametype{
	case "rps":
		var jsonBlob = []byte(`
            {"game": [{"`+gametype+`": {"wins": `+strconv.Itoa(wins)+`, "Losses": `+strconv.Itoa(losses)+`, "Ties": `+strconv.Itoa(ties)+`}}]}
        `)
		rps := Games{}
		err := json.Unmarshal(jsonBlob, &rps)
		if err != nil {
			fmt.Println("Error")
		}
		saveJson, _ := json.Marshal(rps)
		err = ioutil.WriteFile(gametype+".json", saveJson, 0644)
		fmt.Println("%+v", rps)
	case "blackjack":
		var jsonBlob = []byte(`
            {"game": [{"`+gametype+`": {"wins": `+strconv.Itoa(wins)+`, "Losses": `+strconv.Itoa(losses)+`, "Ties": `+strconv.Itoa(ties)+`}}]}
        `)
		bj := Games2{}
		err := json.Unmarshal(jsonBlob, &bj)
		if err != nil {
			fmt.Println("Error")
		}
		saveJson, _ := json.Marshal(bj)
		err = ioutil.WriteFile(gametype+".json", saveJson, 0644)
		fmt.Println("%+v", rps)
	}
}

func main() {
	dg, err := discordgo.New("Bot " + "oops")
	if err != nil {
		fmt.Println("error created while making a bot")
		return
	}
	dg.AddHandler(on_message)
	dg.AddHandler(rps)
	dg.AddHandler(help)
	dg.AddHandler(ping)
	dg.AddHandler(source)
	dg.AddHandler(stats)
	dg.AddHandler(black_jack)
	dg.AddHandler(invite)
	dg.AddHandler(on_ready)
	err = dg.Open()
	if err != nil {
		fmt.Println("Error created while opening the bot", err)
		return
	}
	fmt.Println("Bot is up and running :sunglasses:")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

}
func on_ready(s *discordgo.Session, e *discordgo.Ready) {
	fmt.Println("Bot online")
	s.UpdateGameStatus(0, ".help | Made in golang with love")
}
func EmbedMsgHello(game string, title string, description string, Thumbnail string) *discordgo.MessageEmbed {
	switch game {
	case "bj":
		embed := & discordgo.MessageEmbed {
			Fields: [] * discordgo.MessageEmbedField { & discordgo.MessageEmbedField {
				Name: "*",
				Value: description,
				Inline: true,
			},
			},
			Thumbnail: & discordgo.MessageEmbedThumbnail {
				URL: Thumbnail,
			},
			Footer: & discordgo.MessageEmbedFooter {
				Text: "Made by Leki#6796",
			},
			Timestamp: time.Now().Format(time.RFC3339),
			Title: title,
		}
		return embed
	case "rps":
		embed := & discordgo.MessageEmbed {
			Fields: [] * discordgo.MessageEmbedField { & discordgo.MessageEmbedField {
				Name: "*",
				Value: description,
				Inline: true,
			},
			},
			Image: & discordgo.MessageEmbedImage {
				URL: Thumbnail,
			},
			Footer: & discordgo.MessageEmbedFooter {
				Text: "Made by Leki#6796",
			},
			Timestamp: time.Now().Format(time.RFC3339),
			Title: title,
		}
		return embed
	}
	embed := & discordgo.MessageEmbed {
		Fields: [] * discordgo.MessageEmbedField { & discordgo.MessageEmbedField {
			Name: "Missing information",
			Value: "The developer of this bot is missing some information for this function, hopefully his dumbass can fix it soon",
			Inline: true,
		},
		},
		Thumbnail: & discordgo.MessageEmbedThumbnail {
			URL: "https://i.imgur.com/NldSwaZ.png",
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Title: "Golang Game Bot",
	}
	return embed
}
// ----------INVITE COMMAND ---------\\
func invite(s * discordgo.Session, m* discordgo.MessageCreate) {
	if m.Content == ".invite" {
		if m.Author.ID != s.State.User.ID {
			s.ChannelMessageSendEmbed(m.ChannelID, EmbedMsgHello("bj", "Invite me to your server!", "[Click here!](https://discord.com/api/oauth2/authorize?client_id=905891689581916210&permissions=314432&scope=bot)", "https://i.imgur.com/NldSwaZ.png"))
		}
	}
}
// ----------END INVITE COMMAND ---------\\

// --------HELP COMMAND------------\\
func help(s * discordgo.Session, m * discordgo.MessageCreate) {
	if m.Content == ".help" {
		if m.Author.ID != s.State.User.ID {

			s.ChannelMessageSendEmbed(m.ChannelID, EmbedMsgHello("bj", "Help Commands", ".ping\n.bj\n.rps\n.source", "https://i.imgur.com/NldSwaZ.png"))

		}
	}
}
// --------END HELP COMMAND------------\\
// --------BLACK JACK COMMAND------------\\
func black_jack(s * discordgo.Session, m * discordgo.MessageCreate) {
	if m.Content == ".bj" {
		originaluserid = m.Author.ID
		msg, err := s.ChannelMessageSendEmbed(m.ChannelID, EmbedMsgHello("bj", "Discord Black Jack", "React below to start the game!", "https://imgur.com/BbgsSmC.png"))
		if err != nil {
			fmt.Println(err)
		}
		s.MessageReactionAdd(m.ChannelID, msg.ID, "🎴")
		originalmessageid = msg.ID
	}
}
// --------END BLACK JACK COMMAND------------\\
// --------------------STATS COMMAND------------\\
func stats(s * discordgo.Session, m * discordgo.MessageCreate) {
	if m.Content == ".stats" {
		if m.Author.ID != s.State.User.ID {
			jsonFile, err := os.Open("rps.json")
			if err != nil {
				fmt.Println(err)
			}
			var gamer Games
			fmt.Println("Successfully Opened users.json")
			defer jsonFile.Close()
			byteValue, _ := ioutil.ReadAll(jsonFile)

			json.Unmarshal(byteValue, &gamer)
			s.ChannelMessageSendEmbed(m.ChannelID, EmbedMsgHello("bj", "Rock Paper Scissors Total game stats", "Wins: " + strconv.Itoa(gamer.Games[0].Rps.Wins) + "\nLosses: " + strconv.Itoa(gamer.Games[0].Rps.Losses) + "\nTies: " + strconv.Itoa(gamer.Games[0].Rps.Ties), "https://i.imgur.com/i7G2Yuc.png"))
			jsonfile, err := os.Open("blackjack.json")
			if err!= nil {
				fmt.Println("Error", err)
			}
			var bj Games2
			fmt.Println("Opended file")
			defer jsonfile.Close()
			biteValue, _ := ioutil.ReadAll(jsonfile)
			json.Unmarshal(biteValue, &bj)
			s.ChannelMessageSendEmbed(m.ChannelID, EmbedMsgHello("bj", "Black Jack Total game stats", "Wins: " + strconv.Itoa(bj.Games2[0].BlackJack.Wins) + "\nLosses: " + strconv.Itoa(bj.Games2[0].BlackJack.Losses) + "\nTies: " + strconv.Itoa(bj.Games2[0].BlackJack.Ties), "https://imgur.com/BbgsSmC.png"))
		}
	}
}
// --------------------END STATS COMMAND------------\\
// --------------------SOURCE COMMAND------------\\
func source(s * discordgo.Session, m * discordgo.MessageCreate) {
	if m.Author.ID != s.State.User.ID {
		if m.Content == ".source" {
			s.ChannelMessageSendEmbed(m.ChannelID, EmbedMsgHello("bj", "Source code!", "Visit the repository here, https://github.com/zLeki/discord.go-Golang-Game-Bot", "https://opengraph.githubassets.com/1d74613fd71a8b3a3271d6bcf224e2f382c09ceaa7a67da525473e2658fbfea2/zLeki/discord.go-Golang-Game-Bot"))
		}
	}
}
// --------------------END SOURCE COMMAND------------\\
// --------------------PING COMMAND------------\\
func ping(s * discordgo.Session, m * discordgo.MessageCreate) {
	if m.Content == ".ping" {
		if m.Author.ID != s.State.User.ID {
			s.ChannelMessageSend(m.ChannelID, s.HeartbeatLatency().String())
			return
		}
	}
}
// --------------------ROCK PAPER SCISSORS------------\\
func on_message(s * discordgo.Session, m * discordgo.MessageCreate) {

	if m.Author.ID != s.State.User.ID {
		if m.Content == ".rps" {
			originaluserid = m.Author.ID

			msg, err := s.ChannelMessageSendEmbed(m.ChannelID, EmbedMsgHello("bj", "Choose a option", "Rock, Paper, or Scissors", "https://i.imgur.com/gFFSVll.png"))
			if err != nil {
				fmt.Println("error when sending message", err)
				return
			}
			originalmessageid = msg.ID
			s.MessageReactionAdd(m.ChannelID, msg.ID, "🗿")
			s.MessageReactionAdd(m.ChannelID, msg.ID, "📄")
			s.MessageReactionAdd(m.ChannelID, msg.ID, "✂️")
		}
	}
}
// --------END ROCK PAPER SCISSORS COMMAND------------\\


// -------RPS FUNCTIONS------------\\
func rps(s * discordgo.Session, r * discordgo.MessageReactionAdd) {
	if r.UserID != s.State.User.ID {
		if r.UserID == originaluserid {
			fmt.Println(r.Emoji.Name, r.ChannelID)
			player1hand := []int{}
			player2hand := []int{}
			var wins = 0
			var losses = 0
			var ties = 0
			fmt.Println(wins, losses, ties)
			savetoJson("rps", wins, losses, ties)
			reasons := make([]string, 0)
			var paper = ":page_facing_up:"
			reasons = append(reasons,
				"rock",  // rock
				"paper", // paper
				"scissors") // scissors
			selected := reasons[rand.Intn(len(reasons))]

			if r.Emoji.Name == "🎴" { // BlackJack part 1
				for i := 0;
					i < 2;
				i++ {
					min := 2
					max := 11
					rand.Seed(time.Now().UnixNano())
					player1hand = append(player1hand, rand.Intn(max-min)+min)
				}
				for i := 0;
					i < 2;
				i++ {
					min := 2
					max := 11
					rand.Seed(time.Now().UnixNano())
					player2hand = append(player2hand, rand.Intn(max-min)+min)
				}
				if player1hand[0]+player1hand[1] != 21 || player2hand[0]+player2hand[1] != 21 {
					enemytotal = player2hand[0] + player2hand[1]
					playertotal = player1hand[0] + player1hand[1]
					embed := &discordgo.MessageEmbed{
						Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{
							Name:   "Black jack",
							Value:  "Opponent hand \n**" + strconv.Itoa(player2hand[0]) + "+" + strconv.Itoa(player2hand[1]) + " = " + strconv.Itoa(enemytotal) + "**\n\nYour hand: \n**" + strconv.Itoa(player1hand[0]) + "+" + strconv.Itoa(player1hand[1]) + " = " + strconv.Itoa(playertotal) + "**",
							Inline: true,
						},
						},
						Thumbnail: &discordgo.MessageEmbedThumbnail{
							URL: "https://imgur.com/BbgsSmC.png",
						},
						Timestamp: time.Now().Format(time.RFC3339),
						Title:     "Discord black jack",
					}
					msg, err := s.ChannelMessageSendEmbed(r.ChannelID, embed)

					fmt.Println(playertotal, enemytotal)
					if err != nil {
						fmt.Println(err)
					}
					originalmessageid = msg.ID
					s.MessageReactionAdd(r.ChannelID, msg.ID, "⬆️")
					s.MessageReactionAdd(r.ChannelID, msg.ID, "⏸️")

				} else {
					fmt.Println("You lose")
				}
			}
			if r.Emoji.Name == "✂️" {
				fmt.Println("hi")
				if selected != "scissors" {
					if selected == "rock" {
						losses += 1
						savetoJson("rps", wins, losses, ties)
						fmt.Println(wins, losses, ties)
						s.ChannelMessageSendEmbed(r.ChannelID, EmbedMsgHello("rps", "You lose!", "You lose :skull: "+r.Emoji.Name+":left_facing_fist:"+":"+selected+":",
							"https://gyazo.com/3d4362903ae98dce9d36898f45cff353.gif"))
					} else if selected == "paper" {
						s.ChannelMessageSendEmbed(r.ChannelID, EmbedMsgHello("rps", "You win!", "You win :tada: "+r.Emoji.Name+":right_facing_fist:"+paper, "https://gyazo.com/8d200691ec87e5857708ada94c978d2c.gif"))
						wins += 1
						savetoJson("rps", wins, losses, ties)
					}
				} else {
					s.ChannelMessageSendEmbed(r.ChannelID, EmbedMsgHello("rps", "Its a tie!", `"Fair trade" - Drake `+r.Emoji.Name+":handshake:"+":"+selected+":", "https://gyazo.com/6a70a7536cb12cc32cdbfdf16b3942b3.gif"))
					ties += 1
					savetoJson("rps", wins, losses, ties)
				}
			} else if r.Emoji.Name == "📄" {
				if selected != "paper" {
					if selected == "scissors" {
						losses += 1
						savetoJson("rps", wins, losses, ties)
						s.ChannelMessageSendEmbed(r.ChannelID, EmbedMsgHello("rps", "You lose!", "You lose :skull:"+r.Emoji.Name+":left_facing_fist:"+":"+selected+":", "https://gyazo.com/8d200691ec87e5857708ada94c978d2c.gif"))
					} else if selected == "rock" {
						s.ChannelMessageSendEmbed(r.ChannelID, EmbedMsgHello("rps", "You win!", "You win! :tada:"+r.Emoji.Name+":right_facing_fist:"+":"+selected+":", "https://gyazo.com/b451ed25527ea5272274145889bae8a8.gif"))
						wins += 1
						savetoJson("rps", wins, losses, ties)

					}
				} else {
					s.ChannelMessageSendEmbed(r.ChannelID, EmbedMsgHello("rps", "Its a tie!", `"Fair trade" - Drake\n`+r.Emoji.Name+":handshake:"+paper, "https://gyazo.com/19c1084097aec53ec86d2a88627f78a7.gif"))
					ties += 1
					savetoJson("rps", wins, losses, ties)
				}
			} else if r.Emoji.Name == "🗿" {
				fmt.Println(`rock`)
				if selected != "rock" {
					if selected == "paper" {
						losses += 1
						savetoJson("rps", wins, losses, ties)
						s.ChannelMessageSendEmbed(r.ChannelID, EmbedMsgHello("rps", "You lose!", "You lose :skull:\n"+r.Emoji.Name+":left_facing_fist:"+paper, "https://gyazo.com/30f7ca6750c9a9eca2de6dab963591e7.gif"))
					} else if selected == "scissors" {
						s.ChannelMessageSendEmbed(r.ChannelID, EmbedMsgHello("rps", "You win!", "You win :tada:\n"+r.Emoji.Name+":right_facing_fist:"+":"+selected+":", "https://gyazo.com/ae13b787789add0b354bbe14e49d75d4.gif"))
						wins += 1
						savetoJson("rps", wins, losses, ties)
					}
				} else {
					s.ChannelMessageSendEmbed(r.ChannelID, EmbedMsgHello("rps", "Its a tie!", `"Fair trade" - Drake\n`+r.Emoji.Name+":handshake:"+":moyai:", "https://gyazo.com/0fc4cd0690ef3b133bc6a5b0539d03d2.gif"))
					ties += 1
					savetoJson("rps", wins, losses, ties)
				}
				originaluserid = ""
				fmt.Println(originaluserid)
				fmt.Println("tes")
			}
			var bjwins = 0
			var bjlosses = 0
			var bjties = 0
			if r.Emoji.Name == "⏸️" {
				if r.UserID == s.State.User.ID {
					return
				}
				print("hi")
				if enemytotal <= 18 {
					min := 2
					max := 11
					rand.Seed(time.Now().UnixNano())
					enemytotal += rand.Intn(max-min) + min
					s.ChannelMessageSend(r.ChannelID, "**Player2 is hitting**")
					if playertotal <= 21 || enemytotal <= 21 {
						msg, err := s.ChannelMessageSendEmbed(r.ChannelID, EmbedMsgHello("bj", "Black Jack", "Opponent hand \n** = "+strconv.Itoa(enemytotal)+"**\n\nYour hand: \n** = "+strconv.Itoa(playertotal)+"**", "https://imgur.com/BbgsSmC.png"))
						if err != nil {
							fmt.Println(err)
						}
						fmt.Println(playertotal, enemytotal)
						originalmessageid = msg.ID
						s.MessageReactionAdd(r.ChannelID, msg.ID, "⬆️")
						s.MessageReactionAdd(r.ChannelID, msg.ID, "⏸️")
						if enemytotal > 21 || playertotal == 21 {
							bjwins += 1
							s.ChannelMessageSendEmbed(r.ChannelID, EmbedMsgHello("bj", "You win!", "You have "+strconv.Itoa(playertotal*7)+" social credit", "https://i.ytimg.com/vi/7Kl_lU_ZaBw/maxresdefault.jpg"))
							savetoJson("blackjack", bjwins, bjlosses, bjties)
							return
						}
						if playertotal > 21 {
							bjlosses += 1
							s.ChannelMessageSendEmbed(r.ChannelID, EmbedMsgHello("bj", "You lose!", "You have -"+strconv.Itoa(playertotal*7)+" social credit GAMBLING IS NOT ALOUD ON CHINESE LANDS", "https://img.ifunny.co/images/3c4cfb8f47538c3a4be0693fbac0113be89cff56269a20674aa2fc2438108d00_3.jpg"))
							return
							if enemytotal == 21 {
								bjlosses += 1
								s.ChannelMessageSendEmbed(r.ChannelID, EmbedMsgHello("bj", "You lose!", "You have -"+strconv.Itoa(playertotal*7)+" social credit GAMBLING IS NOT ALOUD ON CHINESE LANDS", "https://img.ifunny.co/images/3c4cfb8f47538c3a4be0693fbac0113be89cff56269a20674aa2fc2438108d00_3.jpg"))

								return
							}
						}
					}
				} else {
					s.ChannelMessageSend(r.ChannelID, "**Player2 is staying**")
					if playertotal > enemytotal {
						s.ChannelMessageSendEmbed(r.ChannelID, EmbedMsgHello("bj", "You win!", "You have "+strconv.Itoa(playertotal*7)+" social credit", "https://i.ytimg.com/vi/7Kl_lU_ZaBw/maxresdefault.jpg"))
						bjwins += 1
						savetoJson("blackjack", bjwins, bjlosses, bjties)
						return
					}
					if enemytotal > playertotal {
						bjlosses += 1
						s.ChannelMessageSendEmbed(r.ChannelID, EmbedMsgHello("bj", "You lose!", "You have -"+strconv.Itoa(playertotal*7)+" social credit GAMBLING IS NOT ALOUD ON CHINESE LANDS", "https://img.ifunny.co/images/3c4cfb8f47538c3a4be0693fbac0113be89cff56269a20674aa2fc2438108d00_3.jpg"))
						return
					}
					if enemytotal == playertotal {
						bjties += 1
						savetoJson("blackjack", bjwins, bjlosses, bjties)
						s.ChannelMessageSendEmbed(r.ChannelID, EmbedMsgHello("bj", "Its a tie!", "You have +"+strconv.Itoa(playertotal)+" social credit", "https://wompampsupport.azureedge.net/fetchimage?siteId=7575&v=2&jpgQuality=100&width=700&url=https%3A%2F%2Fi.kym-cdn.com%2Fentries%2Ficons%2Foriginal%2F000%2F027%2F195%2Fcover10.jpg"))
						return
					}
					if playertotal <= 21 || enemytotal <= 21 {
						msg, err := s.ChannelMessageSendEmbed(r.ChannelID, EmbedMsgHello("bj", "Black Jack", "Opponent hand \n** = "+strconv.Itoa(enemytotal)+"**\n\nYour hand: \n** = "+strconv.Itoa(playertotal)+"**", "https://imgur.com/BbgsSmC.png"))
						if err != nil {
							fmt.Println(err)
						}
						fmt.Println(playertotal, enemytotal)
						originalmessageid = msg.ID
						s.MessageReactionAdd(r.ChannelID, msg.ID, "⬆️")
						s.MessageReactionAdd(r.ChannelID, msg.ID, "⏸️")
						if enemytotal > 21 || playertotal == 21 {
							bjwins += 1
							savetoJson("blackjack", bjwins, bjlosses, bjties)
							s.ChannelMessageSendEmbed(r.ChannelID, EmbedMsgHello("bj", "You win!", "You have "+strconv.Itoa(playertotal*7)+" social credit", "https://i.ytimg.com/vi/7Kl_lU_ZaBw/maxresdefault.jpg"))
							return
						}
						if playertotal > 21 {
							bjlosses += 1
							savetoJson("blackjack", bjwins, bjlosses, bjties)
							s.ChannelMessageSendEmbed(r.ChannelID, EmbedMsgHello("bj", "You lose!", "You have -"+strconv.Itoa(playertotal*7)+" social credit GAMBLING IS NOT ALOUD ON CHINESE LANDS", "https://img.ifunny.co/images/3c4cfb8f47538c3a4be0693fbac0113be89cff56269a20674aa2fc2438108d00_3.jpg"))
							return
							if enemytotal == 21 {
								bjlosses += 1
								savetoJson("blackjack", bjwins, bjlosses, bjties)
								s.ChannelMessageSendEmbed(r.ChannelID, EmbedMsgHello("bj", "You lose!", "You have -"+strconv.Itoa(playertotal*7)+" social credit GAMBLING IS NOT ALOUD ON CHINESE LANDS", "https://img.ifunny.co/images/3c4cfb8f47538c3a4be0693fbac0113be89cff56269a20674aa2fc2438108d00_3.jpg"))
								return
							}
						}
					}
				}
			}

			if r.Emoji.Name == "⬆️" {
				if r.UserID == s.State.User.ID {
					return
				}
				if enemytotal <= 18 {
					min := 2
					max := 11
					rand.Seed(time.Now().UnixNano())
					enemytotal += rand.Intn(max-min) + min
					rand.Seed(time.Now().UnixNano())
					playertotal += rand.Intn(max-min) + min
					s.ChannelMessageSend(r.ChannelID, "**Player2 is hitting**")
					if playertotal <= 21 || enemytotal <= 21 {
						msg, err := s.ChannelMessageSendEmbed(r.ChannelID, EmbedMsgHello("bj", "Black Jack", "Opponent hand \n** = "+strconv.Itoa(enemytotal)+"**\n\nYour hand: \n** = "+strconv.Itoa(playertotal)+"**", "https://imgur.com/BbgsSmC.png"))
						if err != nil {
							fmt.Println(err)
						}
						originalmessageid = msg.ID
						s.MessageReactionAdd(r.ChannelID, msg.ID, "⬆️")
						s.MessageReactionAdd(r.ChannelID, msg.ID, "⏸️")
						if enemytotal > 21 || playertotal == 21 {
							bjwins += 1
							savetoJson("blackjack", bjwins, bjlosses, bjties)
							s.ChannelMessageSendEmbed(r.ChannelID, EmbedMsgHello("bj", "You win!", "You have "+strconv.Itoa(playertotal*7)+" social credit", "https://i.ytimg.com/vi/7Kl_lU_ZaBw/maxresdefault.jpg"))
						}
						if playertotal > 21 {
							bjlosses += 1
							savetoJson("blackjack", bjwins, bjlosses, bjties)
							s.ChannelMessageSendEmbed(r.ChannelID, EmbedMsgHello("bj", "You lose!", "You have -"+strconv.Itoa(playertotal*7)+" social credit GAMBLING IS NOT ALOUD ON CHINESE LANDS", "https://img.ifunny.co/images/3c4cfb8f47538c3a4be0693fbac0113be89cff56269a20674aa2fc2438108d00_3.jpg"))

							if enemytotal == 21 {
								bjlosses += 1
								savetoJson("blackjack", bjwins, bjlosses, bjties)
								s.ChannelMessageSendEmbed(r.ChannelID, EmbedMsgHello("bj", "You lose!", "You have -"+strconv.Itoa(playertotal*7)+" social credit GAMBLING IS NOT ALOUD ON CHINESE LANDS", "https://img.ifunny.co/images/3c4cfb8f47538c3a4be0693fbac0113be89cff56269a20674aa2fc2438108d00_3.jpg"))
							}
						}
					}
				} else {
					s.ChannelMessageSend(r.ChannelID, "**Player2 is staying**")
					min := 2
					max := 11
					rand.Seed(time.Now().UnixNano())
					playertotal += rand.Intn(max-min) + min
					if playertotal <= 21 || enemytotal <= 21 {
						msg, err := s.ChannelMessageSendEmbed(r.ChannelID, EmbedMsgHello("bj", "Black Jack", "Opponent hand \n** = "+strconv.Itoa(enemytotal)+"**\n\nYour hand: \n** = "+strconv.Itoa(playertotal)+"**", "https://imgur.com/BbgsSmC.png"))
						if err != nil {
							fmt.Println(err)
						}
						if playertotal == enemytotal {
							s.ChannelMessageSendEmbed(r.ChannelID, EmbedMsgHello("bj", "Its a tie!", "You have +"+strconv.Itoa(playertotal)+" social credit", "https://wompampsupport.azureedge.net/fetchimage?siteId=7575&v=2&jpgQuality=100&width=700&url=https%3A%2F%2Fi.kym-cdn.com%2Fentries%2Ficons%2Foriginal%2F000%2F027%2F195%2Fcover10.jpg"))
							bjties += 1
							savetoJson("blackjack", bjwins, bjlosses, bjties)
							return
						}
						fmt.Println(playertotal, enemytotal)
						originalmessageid = msg.ID
						s.MessageReactionAdd(r.ChannelID, msg.ID, "⬆️")
						s.MessageReactionAdd(r.ChannelID, msg.ID, "⏸️")
						if enemytotal > 21 || playertotal == 21 {
							bjwins += 1
							s.ChannelMessageSendEmbed(r.ChannelID, EmbedMsgHello("bj", "You win!", "You have "+strconv.Itoa(playertotal*7)+" social credit", "https://i.ytimg.com/vi/7Kl_lU_ZaBw/maxresdefault.jpg"))
						}
						if playertotal > 21 {
							bjlosses += 1
							savetoJson("blackjack", bjwins, bjlosses, bjties)
							s.ChannelMessageSendEmbed(r.ChannelID, EmbedMsgHello("bj", "You lose!", "You have -"+strconv.Itoa(playertotal*7)+" social credit GAMBLING IS NOT ALOUD ON CHINESE LANDS", "https://img.ifunny.co/images/3c4cfb8f47538c3a4be0693fbac0113be89cff56269a20674aa2fc2438108d00_3.jpg"))
							if enemytotal == 21 {
								bjlosses += 1
								s.ChannelMessageSendEmbed(r.ChannelID, EmbedMsgHello("bj", "You lose!", "You have -"+strconv.Itoa(playertotal*7)+" social credit GAMBLING IS NOT ALOUD ON CHINESE LANDS", "https://img.ifunny.co/images/3c4cfb8f47538c3a4be0693fbac0113be89cff56269a20674aa2fc2438108d00_3.jpg"))
							}
						}
					}
				}
			}
    }
  }
}
