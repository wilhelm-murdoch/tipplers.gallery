import { dev } from "$app/environment"

export const getImageUrl = (path: string, variant: string) => {
  return `https://cdn.tipplers.gallery/${path}/${variant}`
}

export const getBaseUrl = () => {
  return dev ? "http://localhost:5173" : "https://tipplers.gallery"
}