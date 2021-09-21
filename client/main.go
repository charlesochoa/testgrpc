package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"text/template"

	pb "testgrpc/notification"

	"cloud.google.com/go/pubsub"
	"google.golang.org/grpc"
	"tawesoft.co.uk/go/dialog"
)

type RadioButton struct {
	Name       string
	Value      string
	IsDisabled bool
	IsChecked  bool
	Text       string
}

type PageVariables struct {
	PageTitle        string
	PageRadioButtons []RadioButton
	Answer           string
}

var conn *grpc.ClientConn
var serviceClient pb.NotificationServiceClient

func DisplayRadioButtons(w http.ResponseWriter, r *http.Request) {
	// Display some radio buttons to the user
	fmt.Println("Display")
	Title := "Cual Prefieres?"
	MyRadioButtons := []RadioButton{
		{"animalselect", "perros", false, false, "Perros"},
		{"animalselect", "gatos", false, false, "Gatos"},
	}

	MyPageVariables := PageVariables{
		PageTitle:        Title,
		PageRadioButtons: MyRadioButtons,
		Answer:           "",
	}

	t, err := template.ParseFiles("../views/page.html") //parse the html file homepage.html
	if err != nil {                                     // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                     // if there is an error
		log.Print("template executing error: ", err) //log it
	}

}

func UserSelected(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// r.Form is now either
	// map[animalselect:[cats]] OR
	// map[animalselect:[dogs]]
	// so get the animal which has been selected
	youranimal := r.Form.Get("animalselect")
	MyRadioButtons := []RadioButton{
		{"animalselect", "perros", false, false, "Perros"},
		{"animalselect", "gatos", false, false, "Gatos"},
	}

	SendAnswer("Preferencia: " + youranimal)
	Title := "Has elegido: "
	MyPageVariables := PageVariables{
		PageTitle:        Title,
		PageRadioButtons: MyRadioButtons,
		Answer:           youranimal,
	}

	// generate page by passing page variables into template
	t, err := template.ParseFiles("../views/page.html") //parse the html file homepage.html
	if err != nil {                                     // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                     // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

func generateId() string {
	return "1234567"
}

func pubSubReceiver(projectID string, subscriptionID string) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		panic("cannot register new client >>" + err.Error())
	}
	sub := client.Subscription(subscriptionID)
	err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		s := string(m.Data)
		fmt.Println("m.ID >>", m.ID, "m.content >>", s)

		dialog.Alert("Message from PubSub:\nm.ID >>" + m.ID + "\nm.content >>" + s)
		// TODO: Handle message.
		// NOTE: May be called concurrently; synchronize access to shared memory.
		m.Ack()
	})
	if err != context.Canceled {
		panic("context Canceled >> " + err.Error())
	}

}

func SendAnswer(msg string) {

	res, err := serviceClient.Send(context.Background(), &pb.SendItemReq{
		Item: &pb.Item{
			Id:      generateId(),
			Content: msg,
		},
	})
	if err != nil {
		panic("cannot send >> " + err.Error())
	}
	fmt.Println("Message sent: ", res)
}

func main() {
	projectID := "project-prometeo-v2"
	subscriptionID := "test-grpc-subscription"
	go pubSubReceiver(projectID, subscriptionID)
	conn, err := grpc.Dial("testgrpc-web-wibdbssi3q-ew.a.run.app:50051", grpc.WithInsecure())

	if err != nil {
		panic("cannot connect with server " + err.Error())
	}
	serviceClient = pb.NewNotificationServiceClient(conn)

	http.HandleFunc("/", DisplayRadioButtons)
	http.HandleFunc("/selected", UserSelected)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
