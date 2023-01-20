export async function load() {
  const cocktails = await import("../../../../lib/data/cocktails-slim.json")
  return {
    cocktails: cocktails.default
  }
}