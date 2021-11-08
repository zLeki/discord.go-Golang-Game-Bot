package main 
import (
  "fmt"
  "github.com/bwmarrin/discordgo"
  "os"
	"os/signal"
	"syscall"
  "math/rand"
  "time"
)

var originaluserid = ""
var originalmessageid = ""
func main() {
     dg, err := discordgo.New("Bot " + "token.......YYQrYQ.fz9qMd1fmYs7-8QCvfWDM4Ey1dU")
    if err != nil {
        fmt.Println("error created while making a bot")
        return
    }
    dg.AddHandler(on_message)
    dg.AddHandler(on_reaction)
    dg.AddHandler(help)
    dg.AddHandler(ping)
    dg.AddHandler(source)
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
func help(s *discordgo.Session, m *discordgo.MessageCreate) {
  if m.Content == ".help" {
  if m.Author.ID != s.State.User.ID {
    s.ChannelMessageSend(m.ChannelID, "```All commands will be here\n.rps\n.ping```")
    
  }}
}
func source(s *discordgo.Session, m *discordgo.MessageCreate) {
  if m.Author.ID != s.State.User.ID {
    if m.Content == ".source" {
      
    }
  }
}
func ping(s *discordgo.Session, m *discordgo.MessageCreate) {
  if m.Content == ".ping" {
  if m.Author.ID != s.State.User.ID {
  s.ChannelMessageSend(m.ChannelID, s.HeartbeatLatency().String())
  return
}}}
func on_message(s *discordgo.Session, m *discordgo.MessageCreate) {
    
    if m.Author.ID != s.State.User.ID {
        if m.Content == ".rps" {
            originaluserid = m.Author.ID
            msg, err := s.ChannelMessageSend(m.ChannelID, "Rock, Paper, Or Scissors?")
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
func on_reaction(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
    if r.UserID != s.State.User.ID {
      fmt.Println(originaluserid)
      if r.UserID == originaluserid {
        if r.MessageReaction.MessageID == originalmessageid {
        reasons := make([]string, 0)
        var paper = ":page_facing_up:"
        reasons = append(reasons,
                "rock", // rock
                "paper", // paper
                "scissors") // scissors
        selected := reasons[rand.Intn(len(reasons))]
        
        if r.Emoji.Name == "✂️" {
            if selected != "scissors" {
                if selected == "rock" { 
                    
                    embed := &discordgo.MessageEmbed{Author:      &discordgo.MessageEmbedAuthor{}, Description: "Game results", Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{Name:   r.Emoji.Name+":left_facing_fist:"+":"+selected+":", Value:  "You lose :skull:", Inline: true,}, }, Image: &discordgo.MessageEmbedImage{URL: "https://gyazo.com/3d4362903ae98dce9d36898f45cff353.gif",}, Timestamp: time.Now().Format(time.RFC3339), Title:     "You lose :skull:",}
                    s.ChannelMessageSendEmbed(r.ChannelID, embed)
                }else if selected == "paper" {
                    embed := &discordgo.MessageEmbed{Author:      &discordgo.MessageEmbedAuthor{}, Description: "Game results", Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{Name:   r.Emoji.Name+":right_facing_fist:"+paper, Value:  "You win :tada:", Inline: true,}, }, Image: &discordgo.MessageEmbedImage{URL: "https://gyazo.com/b451ed25527ea5272274145889bae8a8.gif",}, Timestamp: time.Now().Format(time.RFC3339), Title:     "You win!",}
                    s.ChannelMessageSendEmbed(r.ChannelID, embed)
                }
            }else {
                embed := &discordgo.MessageEmbed{Author:      &discordgo.MessageEmbedAuthor{}, Description: "Game results", Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{Name:   r.Emoji.Name+":handshake:"+":"+selected+":", Value:  `"Fair trade" - Drake`, Inline: true,}, }, Image: &discordgo.MessageEmbedImage{URL: "https://gyazo.com/6a70a7536cb12cc32cdbfdf16b3942b3.gif",}, Timestamp: time.Now().Format(time.RFC3339), Title:     "Its a tie!",}
                s.ChannelMessageSendEmbed(r.ChannelID, embed)
            }
            
        }else if r.Emoji.Name == "📄" {
            if selected != "paper" {
                if selected == "scissors" {
                  embed := &discordgo.MessageEmbed{Author:      &discordgo.MessageEmbedAuthor{}, Description: "Game results", Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{Name:   r.Emoji.Name+":left_facing_fist:"+":"+selected+":", Value:  "You lose :skull:", Inline: true,}, }, Image: &discordgo.MessageEmbedImage{URL: "https://gyazo.com/8d200691ec87e5857708ada94c978d2c.gif",}, Timestamp: time.Now().Format(time.RFC3339), Title:     "You lose :skull:",}
                    s.ChannelMessageSendEmbed(r.ChannelID, embed)

                }else if selected == "rock" {
                  embed := &discordgo.MessageEmbed{Author:      &discordgo.MessageEmbedAuthor{}, Description: "Game results", Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{Name:   r.Emoji.Name+":right_facing_fist:"+":"+selected+":", Value:  "You win :tada:", Inline: true,}, }, Image: &discordgo.MessageEmbedImage{URL: "https://gyazo.com/b451ed25527ea5272274145889bae8a8.gif",}, Timestamp: time.Now().Format(time.RFC3339), Title:     "You win!",}
                    s.ChannelMessageSendEmbed(r.ChannelID, embed)
                }
            }else {
                embed := &discordgo.MessageEmbed{Author:      &discordgo.MessageEmbedAuthor{}, Description: "Game results", Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{Name:   r.Emoji.Name+":handshake:"+paper, Value:  `"Fair trade" - Drake`, Inline: true,}, }, Image: &discordgo.MessageEmbedImage{URL: "https://gyazo.com/19c1084097aec53ec86d2a88627f78a7.gif",}, Timestamp: time.Now().Format(time.RFC3339), Title:     "Its a tie!",}
                s.ChannelMessageSendEmbed(r.ChannelID, embed)
            }
        
        }else if r.Emoji.Name == "🗿" {
            if selected != "rock" {
                if selected == "paper" {
                  embed := &discordgo.MessageEmbed{Author:      &discordgo.MessageEmbedAuthor{}, Description: "Game results", Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{Name:   r.Emoji.Name+":left_facing_fist:"+paper, Value:  "You lose :skull:", Inline: true,}, }, Image: &discordgo.MessageEmbedImage{URL: "https://gyazo.com/30f7ca6750c9a9eca2de6dab963591e7.gif",}, Timestamp: time.Now().Format(time.RFC3339), Title:     "You lose!",}
              
                    s.ChannelMessageSendEmbed(r.ChannelID, embed)
                }else if selected == "scissors" {
                  embed := &discordgo.MessageEmbed{Author:      &discordgo.MessageEmbedAuthor{}, Description: "Game results",Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{Name:   r.Emoji.Name+":right_facing_fist:"+":"+selected+":", Value:  "You win :tada:", Inline: true,}, }, Image: &discordgo.MessageEmbedImage{URL: "https://gyazo.com/ae13b787789add0b354bbe14e49d75d4.gif",}, Timestamp: time.Now().Format(time.RFC3339), Title:     "You win!",}
                s.ChannelMessageSendEmbed(r.ChannelID, embed)
                }
            }else {
              embed := &discordgo.MessageEmbed{Author:      &discordgo.MessageEmbedAuthor{}, Description: "Game results",Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{Name:   r.Emoji.Name+":handshake:"+":moyai:", Value:  `"Fair trade" - Drake`, Inline: true,}, }, Image: &discordgo.MessageEmbedImage{URL: "https://gyazo.com/0fc4cd0690ef3b133bc6a5b0539d03d2.gif",}, Timestamp: time.Now().Format(time.RFC3339), Title:     "Its a tie!",}
                s.ChannelMessageSendEmbed(r.ChannelID, embed)
            }
          originaluserid = ""
          fmt.Println(originaluserid)
          fmt.Println("tes")
          return
        }}
    }
    }
}
