## Getting Started

First, install this node.js project (node.js 18):

```bash
npm install
```



to run the development server for the app:

```bash
#development in docker container
docker compos build #first time only
docker compose up
#if you face any issues with 'permission denied' run 'chmod 777 amos2023ss04-kubernetes-inventory-taker'

#or on your host machine
npm run dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

You can start editing the page by modifying `pages/index.tsx`. The page auto-updates as you edit the file.

[API routes](https://nextjs.org/docs/api-routes/introduction) can be accessed on [http://localhost:3000/api/hello](http://localhost:3000/api/hello). This endpoint can be edited in `pages/api/hello.ts`.

The `pages/api` directory is mapped to `/api/*`. Files in this directory are treated as [API routes](https://nextjs.org/docs/api-routes/introduction) instead of React pages.



to deploy app:
```bash
docker build -t app -f App-Dockerfile .
docker run app
```


