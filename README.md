# Tippler's Gallery
A compendium of the finest cocktail recipes.

## Install Locally
Ensure you're using a recent version of `node`. This project was initially built using `19.3.0`. 

> Before you continue, be warned that this is a large repository. To keep costs down, I do not keep the associated plant images in remote object storage like S3. They are all stored here. So, if you're not comfortable with using over 1.5GB of local block storage to play with this project, you may want to take a pass.

```
git clone https://github.com/wilhelm-murdoch/plantsm.art.git
cd plantsm.art
npm install
npm run dev

> plantsm.art@0.0.1 dev
> vite dev


  VITE v4.0.4  ready in 1250 ms

  ➜  Local:   http://localhost:5173/
  ➜  Network: use --host to expose
  ➜  press h to show help
```

Open a browser window and point it to [http://localhost:5173/](http://localhost:5173/). You should now be able to modify anything under `/src` and the local dev build will perform a live update in your browser.
