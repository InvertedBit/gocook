package components

import (
	"fmt"

	"github.com/invertedbit/gocook/models"
	"maragu.dev/gomponents"
	htmx "maragu.dev/gomponents-htmx"
	"maragu.dev/gomponents/html"
)

func EditButton(displayName string) gomponents.Node {
	return html.Button(
		html.Class("btn btn-xs ms-2 btn-primary display-inline etf-edit-button hidden"),
		gomponents.Text(fmt.Sprintf("Edit %s", displayName)),
		htmx.On("click", "this.closest('.editable-text-field').querySelector('.etf-form').classList.remove('hidden'); this.closest('.editable-text-field').querySelector('.etf-display-container').classList.add('hidden'); this.closest('.editable-text-field').querySelector('input').focus();"),
	)
}

func SaveButton(displayName string) gomponents.Node {
	return html.Button(
		html.Class("btn btn-primary display-inline save-button"),
		gomponents.Text(fmt.Sprintf("Save %s", displayName)),
	)
}

func CancelButton() gomponents.Node {
	return html.Button(
		html.Class("btn btn-secondary display-inline"),
		gomponents.Text("Cancel"),
		htmx.On("click", "event.preventDefault(); this.closest('.editable-text-field').querySelector('.etf-display-container').classList.remove('hidden'); this.closest('.editable-text-field').querySelector('.etf-form').classList.add('hidden'); this.closest('.editable-text-field').querySelector('input').value = this.closest('.editable-text-field').querySelector('.etf-display').innerHTML; return false;"),
	)
}

func EditableTextField(classes string, endpoint string, name string, value string, displayName string) gomponents.Node {
	return html.Div(
		htmx.On("mouseenter", "this.querySelector('.etf-edit-button').classList.remove('hidden');"),
		htmx.On("mouseleave", "this.querySelector('.etf-edit-button').classList.add('hidden');"),
		html.Class("editable-text-field"),
		html.Div(
			html.Class("etf-display-container"),
			html.Span(
				html.Class(fmt.Sprintf("etf-display %s", classes)),
				gomponents.Text(value),
				// htmx.On("click", "alert('Hi!');"),
			),
			EditButton(displayName),
		),

		html.Form(
			htmx.Put(endpoint),
			htmx.Target("closest .editable-text-field"),
			// htmx.Trigger("closest .save-button click"),
			html.Class("etf-form hidden"),

			html.Div(
				html.Class("join"),
				html.Input(
					html.Name(name),
					html.Type("text"),
					html.Value(value),
					html.Class(fmt.Sprintf("input input-bordered display-inline %s", classes)),
				),

				SaveButton(displayName),
				CancelButton(),
			),
		),
	)
}

func EditableTextArea(classes string, endpoint string, name string, value string, displayName string) gomponents.Node {
	containerClasses := "editable-text-field min-h-4"
	if value == "" {
		containerClasses += " rounded-sm border border-dashed border-gray-400 p-2"
	}

	return html.Div(
		htmx.On("mouseenter", "this.querySelector('.etf-edit-button').classList.remove('hidden');"),
		htmx.On("mouseleave", "this.querySelector('.etf-edit-button').classList.add('hidden');"),
		html.Class(containerClasses),
		html.Div(
			html.Class("etf-display-container flex flex-row"),
			html.P(
				html.Class(fmt.Sprintf("etf-display display-inline %s", classes)),
				gomponents.Text(value),
			),
			EditButton(displayName),
		),

		html.Form(
			htmx.Put(endpoint),
			htmx.Target("closest .editable-text-field"),
			htmx.Swap("outerHTML"),
			// htmx.Trigger("closest .save-button click"),
			html.Class("etf-form hidden"),

			html.Textarea(
				html.Name(name),
				html.Class(fmt.Sprintf("textarea textarea-bordered display-inline %s", classes)),
				gomponents.Text(value),
			),

			SaveButton(displayName),
			CancelButton(),
		),
	)
}

func EditableImage(classes string, endpoint string, image *models.Media, displayName string) gomponents.Node {
	containerClasses := "editable-image-field"
	imageEmbed := html.Img()
	if image != nil {
		containerClasses += " rounded-2xl border border-dashed border-4 border-gray-400 p-2 h-40"
		imageEmbed = html.Img(
			html.Src(image.GetURL()),
		)
	}
	return html.Div(
		html.Class(containerClasses),
		imageEmbed,
		html.Form(
			htmx.Encoding("multipart/form-data"),
			htmx.Put(endpoint),
			htmx.Target("closest .editable-image-field"),
			htmx.Swap("outerHTML"),
			htmx.On("htmx:xhr:progress", "htmx.find('closest .editable-image-field').querySelector('.upload-progress').setAttribute('value', evt.detail.loaded / evt.detail.total * 100)"),
			html.Class("eif-form"),
			html.Input(
				html.Type("file"),
				html.Name("image"),
				html.Class("file-input"),
			),
			html.Button(
				gomponents.Text(fmt.Sprintf("Upload %s", displayName)),
			),
			html.Progress(
				html.Class("upload-progress"),
				html.Value("0"),
				html.Max("100"),
			),
		),
	)
}
