export async function load() {
  const cocktails = await import("../../lib/data/cocktails.json")
  return {
    cocktails: cocktails.default
  }
}