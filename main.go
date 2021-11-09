package main 
import (
  "fmt"
  "github.com/bwmarrin/discordgo"
  "os"
	"os/signal"
	"syscall"
  "math/rand"
  "time"
  "strconv"
)


// Social struct which contains a
// list of links
var wins = 0
var losses = 0
var ties = 0
var bjwins = 0
var bjlosses =0 
var bjties = 0
var playertotal = 0
var enemytotal = 0
var originaluserid = ""
var originalmessageid = ""
func main() {
     dg, err := discordgo.New("Bot " + "#queen #wap #wine4life")
    if err != nil {
        fmt.Println("error created while making a bot")
        return
    }
    dg.AddHandler(on_message)
    dg.AddHandler(on_reaction)
    dg.AddHandler(help)
    dg.AddHandler(ping)
    dg.AddHandler(source)
    dg.AddHandler(stats)
    dg.AddHandler(black_jack)
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


func black_jack(s *discordgo.Session, m *discordgo.MessageCreate) {
  
  if m.Content == ".bj" {
      originaluserid = m.Author.ID
      
      embed := &discordgo.MessageEmbed{Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{Name: "Black jack prompt", Value:  "React below to start the game!", Inline: true,}, }, Thumbnail: &discordgo.MessageEmbedThumbnail{URL: "https://imgur.com/BbgsSmC.png"}, Timestamp: time.Now().Format(time.RFC3339), Title:     "Discord Black Jack",}
      msg, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
      if err != nil {
        fmt.Println(err)
      }
      s.MessageReactionAdd(m.ChannelID, msg.ID, "🎴")
      originalmessageid = msg.ID
  }
}






func stats(s* discordgo.Session, m *discordgo.MessageCreate) {
  if m.Content == ".stats" {
  if m.Author.ID != s.State.User.ID {
  
    
    
    
    
    
    print(wins, ties, losses)    
    embed := &discordgo.MessageEmbed{Author:      &discordgo.MessageEmbedAuthor{}, Description: "Total game stats", Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{Name: "Rock Paper Scissors Stats", Value:  "Wins: "+strconv.Itoa(wins)+"\nLosses: "+strconv.Itoa(losses)+"\nTies: "+strconv.Itoa(ties), Inline: true,}, }, Thumbnail: &discordgo.MessageEmbedThumbnail{URL: "https://i.imgur.com/i7G2Yuc.png"}, Timestamp: time.Now().Format(time.RFC3339)} 

        s.ChannelMessageSendEmbed(m.ChannelID, embed)
        embed2 := &discordgo.MessageEmbed{Author:      &discordgo.MessageEmbedAuthor{}, Description: "Total game stats", Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{Name: "Blackjack Stats", Value:  "Wins: "+strconv.Itoa(bjwins)+"\nLosses: "+strconv.Itoa(bjlosses)+"\nTies: "+strconv.Itoa(bjties), Inline: true,}, }, Thumbnail: &discordgo.MessageEmbedThumbnail{URL: "https://imgur.com/BbgsSmC.png"}, Timestamp: time.Now().Format(time.RFC3339)} 
        s.ChannelMessageSendEmbed(m.ChannelID, embed2)
      }
        // fmt.Println("User Type: " + users.Users[i].Type)
        // fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
        // fmt.Println("User Name: " + users.Users[i].Name)
        // fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
    }

}
  
    // }
    //fmt.Println(strconv.Itoa(users.Users[0].wins))
    

func source(s *discordgo.Session, m *discordgo.MessageCreate) {
  if m.Author.ID != s.State.User.ID {

    if m.Content == ".source" {
      
      embed := &discordgo.MessageEmbed{Author:      &discordgo.MessageEmbedAuthor{}, Description: "Golang Game bot source code", Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{Name: "Visit the repository here", Value:  "https://github.com/zLeki/discord.go-Golang-Game-Bot", Inline: true,}, }, Image: &discordgo.MessageEmbedImage{URL: "https://opengraph.githubassets.com/1d74613fd71a8b3a3271d6bcf224e2f382c09ceaa7a67da525473e2658fbfea2/zLeki/discord.go-Golang-Game-Bot",}, Timestamp: time.Now().Format(time.RFC3339), Title:     "Source code!",}
      s.ChannelMessageSendEmbed(m.ChannelID, embed)
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
      if r.UserID == originaluserid {
        fmt.Println(r.Emoji.Name, r.ChannelID)
        player1hand := []int{}
        player2hand := []int{}
        

        
        reasons := make([]string, 0)
        var paper = ":page_facing_up:"
        reasons = append(reasons,
                "rock", // rock
                "paper", // paper
                "scissors") // scissors
        selected := reasons[rand.Intn(len(reasons))]

        if r.Emoji.Name == "🎴" { // BlackJack part 1
          
          //s.ChannelMessageSend(r.ChannelID, strconv.Itoa())
          for i := 0; i < 2; i++ {
              min := 2
              max := 11
              rand.Seed(time.Now().UnixNano())
            player1hand = append(player1hand, rand.Intn(max - min) + min)
          }
          for i := 0; i < 2; i++ {
              min := 2
              max := 11
              rand.Seed(time.Now().UnixNano())
            player2hand = append(player2hand, rand.Intn(max - min) + min)
          }
          if player1hand[0] + player1hand[1] != 21 || player2hand[0] + player2hand[1] != 21 {
            enemytotal = player2hand[0] + player2hand[1]
            playertotal = player1hand[0] + player1hand[1]
            embed := &discordgo.MessageEmbed{Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{Name: "Black jack", Value:  "Opponent hand \n**"+strconv.Itoa(player2hand[0])+"+"+strconv.Itoa(player2hand[1])+" = "+strconv.Itoa(enemytotal)+"**\n\nYour hand: \n**"+strconv.Itoa(player1hand[0])+"+"+strconv.Itoa(player1hand[1])+" = "+strconv.Itoa(playertotal)+"**", Inline: true,}, }, Thumbnail: &discordgo.MessageEmbedThumbnail{URL: "https://imgur.com/BbgsSmC.png"}, Timestamp: time.Now().Format(time.RFC3339), Title:     "Discord black jack",}
            msg, err := s.ChannelMessageSendEmbed(r.ChannelID, embed)

            fmt.Println(playertotal, enemytotal)
            if err != nil {
              fmt.Println(err)
            }
            originalmessageid = msg.ID
            s.MessageReactionAdd(r.ChannelID, msg.ID, "⬆️")
            s.MessageReactionAdd(r.ChannelID, msg.ID, "⏸️")

          }else{
            fmt.Println("You lose")
          }
        }
        if r.Emoji.Name == "✂️" {
            if selected != "scissors" {
                if selected == "rock" { 

                    
                    embed := &discordgo.MessageEmbed{Author:      &discordgo.MessageEmbedAuthor{}, Description: "Game results", Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{Name:   r.Emoji.Name+":left_facing_fist:"+":"+selected+":", Value:  "You lose :skull:", Inline: true,}, }, Image: &discordgo.MessageEmbedImage{URL: "https://gyazo.com/3d4362903ae98dce9d36898f45cff353.gif",}, Timestamp: time.Now().Format(time.RFC3339), Title:     "You lose :skull:",}
                    losses +=1
                    s.ChannelMessageSendEmbed(r.ChannelID, embed)
                }else if selected == "paper" {
                    
                    embed := &discordgo.MessageEmbed{Author:      &discordgo.MessageEmbedAuthor{}, Description: "Game results", Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{Name:   r.Emoji.Name+":right_facing_fist:"+paper, Value:  "You win :tada:", Inline: true,}, }, Image: &discordgo.MessageEmbedImage{URL: "https://gyazo.com/8d200691ec87e5857708ada94c978d2c.gif",}, Timestamp: time.Now().Format(time.RFC3339), Title:     "You win!",}
                    s.ChannelMessageSendEmbed(r.ChannelID, embed)
                    wins+=1
                }
            }else {
                
                embed := &discordgo.MessageEmbed{Author:      &discordgo.MessageEmbedAuthor{}, Description: "Game results", Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{Name:   r.Emoji.Name+":handshake:"+":"+selected+":", Value:  `"Fair trade" - Drake`, Inline: true,}, }, Image: &discordgo.MessageEmbedImage{URL: "https://gyazo.com/6a70a7536cb12cc32cdbfdf16b3942b3.gif",}, Timestamp: time.Now().Format(time.RFC3339), Title:     "Its a tie!",}
                s.ChannelMessageSendEmbed(r.ChannelID, embed)
                ties+=1
            }
            
        }else if r.Emoji.Name == "📄" {
            if selected != "paper" {
                if selected == "scissors" {
                  embed := &discordgo.MessageEmbed{Author:      &discordgo.MessageEmbedAuthor{}, Description: "Game results", Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{Name:   r.Emoji.Name+":left_facing_fist:"+":"+selected+":", Value:  "You lose :skull:", Inline: true,}, }, Image: &discordgo.MessageEmbedImage{URL: "https://gyazo.com/8d200691ec87e5857708ada94c978d2c.gif",}, Timestamp: time.Now().Format(time.RFC3339), Title:     "You lose :skull:",}
                  losses +=1
                    s.ChannelMessageSendEmbed(r.ChannelID, embed)

                }else if selected == "rock" {
                  embed := &discordgo.MessageEmbed{Author:      &discordgo.MessageEmbedAuthor{}, Description: "Game results", Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{Name:   r.Emoji.Name+":right_facing_fist:"+":"+selected+":", Value:  "You win :tada:", Inline: true,}, }, Image: &discordgo.MessageEmbedImage{URL: "https://gyazo.com/b451ed25527ea5272274145889bae8a8.gif",}, Timestamp: time.Now().Format(time.RFC3339), Title:     "You win!",}
                    s.ChannelMessageSendEmbed(r.ChannelID, embed)
                    wins+=1
                }
            }else {
                embed := &discordgo.MessageEmbed{Author:      &discordgo.MessageEmbedAuthor{}, Description: "Game results", Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{Name:   r.Emoji.Name+":handshake:"+paper, Value:  `"Fair trade" - Drake`, Inline: true,}, }, Image: &discordgo.MessageEmbedImage{URL: "https://gyazo.com/19c1084097aec53ec86d2a88627f78a7.gif",}, Timestamp: time.Now().Format(time.RFC3339), Title:     "Its a tie!",}
                s.ChannelMessageSendEmbed(r.ChannelID, embed)
                ties+=1
            }
        
        }else if r.Emoji.Name == "🗿" {
            if selected != "rock" {
                if selected == "paper" {
                  embed := &discordgo.MessageEmbed{Author:      &discordgo.MessageEmbedAuthor{}, Description: "Game results", Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{Name:   r.Emoji.Name+":left_facing_fist:"+paper, Value:  "You lose :skull:", Inline: true,}, }, Image: &discordgo.MessageEmbedImage{URL: "https://gyazo.com/30f7ca6750c9a9eca2de6dab963591e7.gif",}, Timestamp: time.Now().Format(time.RFC3339), Title:     "You lose!",}
              losses +=1
                    s.ChannelMessageSendEmbed(r.ChannelID, embed)
                }else if selected == "scissors" {
                  embed := &discordgo.MessageEmbed{Author:      &discordgo.MessageEmbedAuthor{}, Description: "Game results",Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{Name:   r.Emoji.Name+":right_facing_fist:"+":"+selected+":", Value:  "You win :tada:", Inline: true,}, }, Image: &discordgo.MessageEmbedImage{URL: "https://gyazo.com/ae13b787789add0b354bbe14e49d75d4.gif",}, Timestamp: time.Now().Format(time.RFC3339), Title:     "You win!",}
                s.ChannelMessageSendEmbed(r.ChannelID, embed)
                wins+=1
                }
            }else {
              embed := &discordgo.MessageEmbed{Author:      &discordgo.MessageEmbedAuthor{}, Description: "Game results",Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{Name:   r.Emoji.Name+":handshake:"+":moyai:", Value:  `"Fair trade" - Drake`, Inline: true,}, }, Image: &discordgo.MessageEmbedImage{URL: "https://gyazo.com/0fc4cd0690ef3b133bc6a5b0539d03d2.gif",}, Timestamp: time.Now().Format(time.RFC3339), Title:     "Its a tie!",}
                s.ChannelMessageSendEmbed(r.ChannelID, embed)
              ties+=1
            }
        
          originaluserid = ""
          fmt.Println(originaluserid)
          fmt.Println("tes")
          
          

      }}
    }
 
        if r.Emoji.Name == "⏸️" {
          if r.UserID == s.State.User.ID {
            return
          }
          print("hi")
          if enemytotal <= 18 {
            min := 2
            max := 11
            rand.Seed(time.Now().UnixNano())
            enemytotal += rand.Intn(max - min) + min
            s.ChannelMessageSend(r.ChannelID, "**Player2 is hitting**")
        
            if playertotal <= 21 || enemytotal <= 21 {
            embed := &discordgo.MessageEmbed{Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{Name: "Black jack", Value:  "Opponent hand \n** = "+strconv.Itoa(enemytotal)+"**\n\nYour hand: \n** = "+strconv.Itoa(playertotal)+"**", Inline: true,}, }, Thumbnail: &discordgo.MessageEmbedThumbnail{URL: "https://imgur.com/BbgsSmC.png"}, Timestamp: time.Now().Format(time.RFC3339), Title:     "Discord black jack",}
            msg, err := s.ChannelMessageSendEmbed(r.ChannelID, embed)
            if err != nil {
                fmt.Println(err)
              }
              fmt.Println(playertotal, enemytotal)
            originalmessageid = msg.ID
              s.MessageReactionAdd(r.ChannelID, msg.ID, "⬆️")
              s.MessageReactionAdd(r.ChannelID, msg.ID, "⏸️")
            if enemytotal > 21 || playertotal == 21{
            embed := &discordgo.MessageEmbed{Thumbnail: &discordgo.MessageEmbedThumbnail{URL: "https://i.imgur.com/weMIf70.png"}, Timestamp: time.Now().Format(time.RFC3339), Title:     "You win!",}
            bjwins +=1
            s.ChannelMessageSendEmbed(r.ChannelID, embed)
            
            return
            }
            if playertotal > 21 {
            embed := &discordgo.MessageEmbed{Thumbnail: &discordgo.MessageEmbedThumbnail{URL: "https://i.imgur.com/v3w00UP.png"}, Timestamp: time.Now().Format(time.RFC3339), Title:     "You lose!",}
            bjlosses+=1
            s.ChannelMessageSendEmbed(r.ChannelID, embed)
            
            return
            if enemytotal == 21 {
            embed := &discordgo.MessageEmbed{Thumbnail: &discordgo.MessageEmbedThumbnail{URL: "https://i.imgur.com/v3w00UP.png"}, Timestamp: time.Now().Format(time.RFC3339), Title:     "You lose!",}
            bjlosses+=1
            s.ChannelMessageSendEmbed(r.ChannelID, embed)
            
            return
            }}}
          }else{
          s.ChannelMessageSend(r.ChannelID, "**Player2 is staying**")
          
          if playertotal > enemytotal {
            embed := &discordgo.MessageEmbed{Thumbnail: &discordgo.MessageEmbedThumbnail{URL: "https://i.imgur.com/weMIf70.png"}, Timestamp: time.Now().Format(time.RFC3339), Title:     "You win!",}
            s.ChannelMessageSendEmbed(r.ChannelID, embed)
            bjwins+=1
            return
            
          }
          if enemytotal > playertotal {
                        embed := &discordgo.MessageEmbed{Thumbnail: &discordgo.MessageEmbedThumbnail{URL: "https://i.imgur.com/v3w00UP.png"}, Timestamp: time.Now().Format(time.RFC3339), Title:     "You lose!",}
              bjlosses+=1
            s.ChannelMessageSendEmbed(r.ChannelID, embed)
            return
          }
          if enemytotal == playertotal {
                        embed := &discordgo.MessageEmbed{Thumbnail: &discordgo.MessageEmbedThumbnail{URL: "https://imgur.com/eYfecUy.png"}, Timestamp: time.Now().Format(time.RFC3339), Title:     "Its a tie!",}
                        bjties+=1
            s.ChannelMessageSendEmbed(r.ChannelID, embed)
            return
          }
          if playertotal <= 21 || enemytotal <= 21 {
            embed := &discordgo.MessageEmbed{Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{Name: "Black jack", Value:  "Opponent hand \n** = "+strconv.Itoa(enemytotal)+"**\n\nYour hand: \n** = "+strconv.Itoa(playertotal)+"**", Inline: true,}, }, Thumbnail: &discordgo.MessageEmbedThumbnail{URL: "https://imgur.com/BbgsSmC.png"}, Timestamp: time.Now().Format(time.RFC3339), Title:     "Discord black jack",}
            msg, err := s.ChannelMessageSendEmbed(r.ChannelID, embed)
            if err != nil {
                fmt.Println(err)
              }
              fmt.Println(playertotal, enemytotal)
            originalmessageid = msg.ID
              s.MessageReactionAdd(r.ChannelID, msg.ID, "⬆️")
              s.MessageReactionAdd(r.ChannelID, msg.ID, "⏸️")
            if enemytotal > 21 || playertotal == 21{
            embed := &discordgo.MessageEmbed{Thumbnail: &discordgo.MessageEmbedThumbnail{URL: "https://i.imgur.com/weMIf70.png"}, Timestamp: time.Now().Format(time.RFC3339), Title:     "You win!",}
            bjwins+=1
            s.ChannelMessageSendEmbed(r.ChannelID, embed)
            
            return
            }
            if playertotal > 21 {
            embed := &discordgo.MessageEmbed{Thumbnail: &discordgo.MessageEmbedThumbnail{URL: "https://i.imgur.com/v3w00UP.png"}, Timestamp: time.Now().Format(time.RFC3339), Title:     "You lose!",}
            bjlosses+=1
            s.ChannelMessageSendEmbed(r.ChannelID, embed)
            
            return
            if enemytotal == 21 {
            embed := &discordgo.MessageEmbed{Thumbnail: &discordgo.MessageEmbedThumbnail{URL: "https://i.imgur.com/v3w00UP.png"}, Timestamp: time.Now().Format(time.RFC3339), Title:     "You lose!",}
            bjlosses+=1
            s.ChannelMessageSendEmbed(r.ChannelID, embed)
            
            return
            }}}
            
          }
        }
        if r.Emoji.Name == "⬆️" {
          if r.UserID == s.State.User.ID {
            return
          }
          
          // for i := 0; i < 2; i++ {
          //     min := 2
          //     max := 11
          //   player1hand = append(player1hand, rand.Intn(max - min) + min)
          //}
          
          if enemytotal <= 18 {
            min := 2
            max := 11
            rand.Seed(time.Now().UnixNano())
            enemytotal += rand.Intn(max - min) + min
            rand.Seed(time.Now().UnixNano())
            playertotal += rand.Intn(max - min) + min
            s.ChannelMessageSend(r.ChannelID, "**Player2 is hitting**")
            if playertotal <= 21 || enemytotal <= 21 {
            embed := &discordgo.MessageEmbed{Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{Name: "Black jack", Value:  "Opponent hand \n** = "+strconv.Itoa(enemytotal)+"**\n\nYour hand: \n** = "+strconv.Itoa(playertotal)+"**", Inline: true,}, }, Thumbnail: &discordgo.MessageEmbedThumbnail{URL: "https://imgur.com/BbgsSmC.png"}, Timestamp: time.Now().Format(time.RFC3339), Title:     "Discord black jack",}
            msg, err := s.ChannelMessageSendEmbed(r.ChannelID, embed)
            if err != nil {
                fmt.Println(err)
              }
              fmt.Println(playertotal, enemytotal)
            originalmessageid = msg.ID
              s.MessageReactionAdd(r.ChannelID, msg.ID, "⬆️")
              s.MessageReactionAdd(r.ChannelID, msg.ID, "⏸️")
            if enemytotal > 21 || playertotal == 21{
            embed := &discordgo.MessageEmbed{Thumbnail: &discordgo.MessageEmbedThumbnail{URL: "https://i.imgur.com/weMIf70.png"}, Timestamp: time.Now().Format(time.RFC3339), Title:     "You win!",}
            bjwins+=1
            s.ChannelMessageSendEmbed(r.ChannelID, embed)
            
            
            }
            if playertotal > 21 {
            embed := &discordgo.MessageEmbed{Thumbnail: &discordgo.MessageEmbedThumbnail{URL: "https://i.imgur.com/v3w00UP.png"}, Timestamp: time.Now().Format(time.RFC3339), Title:     "You lose!",}
            bjlosses+=1
            s.ChannelMessageSendEmbed(r.ChannelID, embed)
            
            
            if enemytotal == 21 {
            embed := &discordgo.MessageEmbed{Thumbnail: &discordgo.MessageEmbedThumbnail{URL: "https://i.imgur.com/v3w00UP.png"}, Timestamp: time.Now().Format(time.RFC3339), Title:     "You lose!",}
            bjlosses+=1
            s.ChannelMessageSendEmbed(r.ChannelID, embed)
            
            
            }}}
          }else{
          s.ChannelMessageSend(r.ChannelID, "**Player2 is staying**")
          min := 2
          max := 11
          rand.Seed(time.Now().UnixNano())
          playertotal += rand.Intn(max - min) + min
          if playertotal <= 21 || enemytotal <= 21 {
            embed := &discordgo.MessageEmbed{Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{Name: "Black jack", Value:  "Opponent hand \n** = "+strconv.Itoa(enemytotal)+"**\n\nYour hand: \n** = "+strconv.Itoa(playertotal)+"**", Inline: true,}, }, Thumbnail: &discordgo.MessageEmbedThumbnail{URL: "https://imgur.com/BbgsSmC.png"}, Timestamp: time.Now().Format(time.RFC3339), Title:     "Discord black jack",}
            msg, err := s.ChannelMessageSendEmbed(r.ChannelID, embed)
            if err != nil {
                fmt.Println(err)
              }
              fmt.Println(playertotal, enemytotal)
            originalmessageid = msg.ID
              s.MessageReactionAdd(r.ChannelID, msg.ID, "⬆️")
              s.MessageReactionAdd(r.ChannelID, msg.ID, "⏸️")
            if enemytotal > 21 || playertotal == 21{
            embed := &discordgo.MessageEmbed{Thumbnail: &discordgo.MessageEmbedThumbnail{URL: "https://i.imgur.com/weMIf70.png"}, Timestamp: time.Now().Format(time.RFC3339), Title:     "You win!",}
            bjwins+=1
            s.ChannelMessageSendEmbed(r.ChannelID, embed)
            
            
            }
            if playertotal > 21 {
            embed := &discordgo.MessageEmbed{Thumbnail: &discordgo.MessageEmbedThumbnail{URL: "https://i.imgur.com/v3w00UP.png"}, Timestamp: time.Now().Format(time.RFC3339), Title:     "You lose!",}
            bjlosses+=1
            s.ChannelMessageSendEmbed(r.ChannelID, embed)
            
            
            if enemytotal == 21 {
            embed := &discordgo.MessageEmbed{Thumbnail: &discordgo.MessageEmbedThumbnail{URL: "https://i.imgur.com/v3w00UP.png"}, Timestamp: time.Now().Format(time.RFC3339), Title:     "You lose!",}
            bjlosses+=1
            s.ChannelMessageSendEmbed(r.ChannelID, embed)
            
            
            }}}
          }
          }
        }

