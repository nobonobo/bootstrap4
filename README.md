# Bootstrap4 components for Vecty

- Depends
  - [GopherJS](htts://github.com/gopherjs/gopherjs)
  - [Vecty](https://github.com/gopherjs/vecty)
  - [bootstrap4](https://getbootstrap.com/docs/4.0/getting-started/introduction/)

## Policy

- Support:
  - Primitive tag only.
  - Custom markup attributes for each tag.
  - Custom child components or hml for each tag.

## Component definition pattern

```go
type ComponentName struct {
  vecty.Core
  Bold bool `vecty:"prop"`
  ...
  Markup   vecty.MarkupList      `vecty:"prop"`
  Children vecty.ComponentOrHTML `vecty:"prop"`
}

func (c *ComponentName) Render() vecty.ComponentOrHTML {
  return vecty.Tag(elementName,
    vecty.Markup(
      vecty.ClassMap{
          "bold": c.Bold,
      }
      ...
    ),
    c.Markup,
    c.Children,
  )
}
```

## Usage

### Basic

```go
&ComponentName{
  Bold: true,
  Markup:vecty.Markup(vecty.Attribute("id","name")),
  Children: vecty.Text("name"),
}
```

### Multiple Child

```go
&ComponentName{
  Bold: true,
  Markup:vecty.Markup(vecty.Attribute("id","name")),
  Children: vecty.List(
    vecty.Text("bra..bra.."),
    elem.Span(vecty.Text("memo")),
  ),
}
```

### Nested Component

```go
&ComponentName1{
  Bold: true,
  Markup:vecty.Markup(vecty.Attribute("id","name")),
  Children: &ComponentName2{
    ...
  },
}
```

### Event Handling

logging and original behavier

```go
&ComponentName{
    Href:    "#/",
    Markup: vecty.Markup(event.Click(func(event *vecty.Event) {
        log.Println("click: link")
    })),
    Children: vecty.Text("link"),
},
```

logging only

```go
&ComponentName{
    Href:    "#/",
    Markup: vecty.Markup(event.Click(func(event *vecty.Event) {
        log.Println("click: link")
    }).PreventDefault()),
    Children: vecty.Text("link"),
},
```

## TODO

### Supported Coponents

- [x] Alerts
- [x] Badge
- [ ] Breadcrumb
- [x] Buttons
- [x] Button group
- [x] Card
- [ ] Carousel
- [ ] Collapse
- [x] Dropdowns
- [x] Forms
- [x] Input group
- [x] Jumbotron
- [x] List group
- [x] Modal
- [ ] Navs
- [ ] Navbar
- [ ] Pagination
- [x] Popovers
- [x] Progress
- [ ] Scrollspy
- [x] Tooltips
