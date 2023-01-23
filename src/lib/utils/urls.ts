import { dev, building } from "$app/environment"

export const getImageUrl = (path: string, variant: string) => {
  return `https://cdn.tipplers.gallery/${path}/${variant}`
}

export const getBaseUrl = () => {
  return ""
  return dev || building ? "http://localhost:5173" : "https://tipplers.gallery"
}