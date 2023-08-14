package error

import (
	"context"
	"log"
	"time"
)

//Tips for error handling

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestId", "123")
	go someApiCall(ctx)
	time.Sleep(1)
}

func someApiCall(ctx context.Context) {
	err := someUncertainOperation()
	if err != nil {
		log.Fatalf("failed for requestId={%s} and error = {%s}", ctx.Value("requestId"), err)
	}
}
