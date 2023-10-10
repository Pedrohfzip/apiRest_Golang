package controller
import(
	"net/http"
	entities "api/api/entities"
	"github.com/gin-gonic/gin"
)
type tweetController struct{
	tweets []entities.Tweet
}
func NewTweetController() *tweetController{
	return &tweetController{}
}
func (t *tweetController) FindAll(ctx *gin.Context){
	ctx.JSON(http.StatusOK, t.tweets)
}
func (t *tweetController) Create(ctx *gin.Context){
	tweet := entities.NewTweet()
	if err := ctx.BindJSON(&tweet); err != nil {
		return
	}
	t.tweets = append(t.tweets, *tweet)
	ctx.JSON(http.StatusOK, t.tweets)
}
func (t *tweetController) Delete(ctx *gin.Context){
	id := ctx.Param("id")
	for index, tweet := range t.tweets {
		if tweet.ID == id{
			t.tweets = append(t.tweets[0:index], t.tweets[index+1:]...)
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "tweet not found",
	})
}


