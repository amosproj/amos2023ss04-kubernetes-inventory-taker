import useSWR from 'swr'

const fetcher = (url: RequestInfo | URL) => fetch(url).then((res) => res.json());

export default function Home() {
  const { data, error, isLoading } = useSWR(
    '/api/hello',
    fetcher
  );



  if (error) return <div>failed to load</div>
  if (isLoading) return <div>loading...</div>
  return (
    <>
      <h1 className="text-2xl font-semibold text-[#326ce5] mt-6 text-center">Kubernetes Inventory Taker</h1>
      <p className="text-center">by the way: #326ce5 is the hex code of the kubernetes blue color ...</p>
      <p className="text-center">Fetched data coming from Backend: {data.name}</p>
    </>
  )
}
