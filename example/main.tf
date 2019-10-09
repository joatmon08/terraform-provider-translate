resource "translate_text" "hello" {
  text = file("test.md")
  source_language = "en"
  target_language = "es"
}

output "translated_text" {
  value = translate_text.hello.translated_text
}