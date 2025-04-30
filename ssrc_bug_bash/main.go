package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/remoteconfig"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("key.json"))
	if err != nil {
		log.Fatalf("[Main] Error in creating app : %v ", err)
	}

	rc, rcErr := app.RemoteConfig(ctx)
	if rcErr != nil {
		log.Fatalf("[Main] Error in creating RC client : %v ", rcErr)
	}

	// ************* ONE ****************
	// Operations on an empty template, using GetServerTemplate
	// emptyTemplateOperations(rc, ctx)

	// ************* TWO ****************
	// Get and InitServerTemplate on a non-empty template
	fetchRemoteAndLocalTemplate(rc, ctx)

	// ************* THREE ****************
	// Template evaluation with different conditions
	templateEvaluation(rc, ctx)

}

func emptyTemplateOperations(rc *remoteconfig.Client, ctx context.Context) {
	// empty default config
	defaultConfig := make(map[string]any)
	template, tErr := rc.GetServerTemplate(ctx, defaultConfig)
	if tErr != nil {
		log.Fatalf("[emptyTemplateOperations] Error in fetching template : %v ", tErr)
	}

	// empty context
	context := make(map[string]any)
	config, cErr := template.Evaluate(context)

	jsonifiedTemplate, _ := template.ToJSON()
	fmt.Println("[emptyTemplateOperations] jsonified template : ", jsonifiedTemplate)

	if cErr != nil {
		log.Fatalf("[emptyTemplateOperations] Error in evaluating template : %v ", cErr)
	}
	// config should be empty
	fmt.Println(config)
}

func fetchRemoteAndLocalTemplate(rc *remoteconfig.Client, ctx context.Context) {
	// empty default config
	defaultConfig := make(map[string]any)

	template, tErr := rc.GetServerTemplate(ctx, defaultConfig)
	if tErr != nil {
		log.Fatalf("[fetchRemoteAndLocalTemplate] Error in fetching template : %v ", tErr)
	}

	// validate that you receive valid JSON
	jsonifiedTemplate, _ := template.ToJSON()
	fmt.Println("[fetchRemoteAndLocalTemplate] jsonified template : ", jsonifiedTemplate)

	// modify the value of the jsonified template, try a valid and invalid case
	templateToInit := jsonifiedTemplate
	templateTwo, tTwoErr := rc.InitServerTemplate(defaultConfig, templateToInit)
	if tTwoErr != nil {
		log.Fatalf("[fetchRemoteAndLocalTemplate] Error in fetching template two : %v ", tTwoErr)
	}

	// validate that you receive valid JSON
	jsonifiedTemplateTwo, _ := templateTwo.ToJSON()
	fmt.Println("[fetchRemoteAndLocalTemplate] jsonified template two : ", jsonifiedTemplateTwo)
}

func templateEvaluation(rc *remoteconfig.Client, ctx context.Context) {
	// pass values based on your template
	defaultConfig := make(map[string]any)

	template, tErr := rc.GetServerTemplate(ctx, defaultConfig)
	if tErr != nil {
		log.Fatalf("[templateEvaluation] Error in fetching template : %v ", tErr)
	}

	// pass values based on your template
	context := make(map[string]any)

	config, cErr := template.Evaluate(context)
	if cErr != nil {
		log.Fatalf("[templateEvaluation] Error in evaluating template : %v ", cErr)
	}

	// Repeat this block for all your parameters
	parameter_name := "param_one"
	fmt.Println("Value source : ", config.GetValueSource(parameter_name) == remoteconfig.Remote)
	fmt.Println("as Boolean ", config.GetBoolean(parameter_name))
	fmt.Println("as Float ", config.GetFloat(parameter_name))
	fmt.Println("as Int ", config.GetInt(parameter_name))
	fmt.Println("as String ", config.GetString(parameter_name))
}
