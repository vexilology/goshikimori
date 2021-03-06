package transform

import "fmt"

func ConvertAchievements(id int) string {
  return fmt.Sprintf("achievements?user_id=%d", id)
}

func ConvertAnime(id int, name string) string {
  return fmt.Sprintf("animes/%d/%s", id, name)
}

func ConvertRoles(id int, name string) string {
  return fmt.Sprintf("%s/%d/roles", name, id)
}

func ConvertSimilar(id int, name string) string {
  return fmt.Sprintf("%s/%d/similar", name, id)
}

func ConvertRelated(id int, name string) string {
  return fmt.Sprintf("%s/%d/related", name, id)
}

func ConvertFranchise(id int, name string) string {
  return fmt.Sprintf("%s/%d/franchise", name, id)
}

func ConvertCalendar(name string) string {
  return fmt.Sprintf("calendar?%s", name)
}

func ConvertUser(id int, name string) string {
  return fmt.Sprintf("users/%d/%s", id, name)
}

func ConvertExternalLinks(id int, name string) string {
  return fmt.Sprintf("%s/%d/external_links", name, id)
}

func ConvertUserRates(id int, name, options string) string {
  return fmt.Sprintf("users/%d/%s?%s", id, name, options)
}
