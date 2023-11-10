import { createResource } from 'solid-js'

export default function App() {
  const fetchMessage = async () => {
    const res = await fetch(`http://localhost:3010/facts`)
    console.log(res)
    return await res.json()
  }
  const [test] = createResource(fetchMessage)
  
  return (
    <h1 class="text-3xl font-bold underline">
      {test.loading && 'Loading...'}
      {test.error && 'Error!'}
      {test() && test().message}
    </h1>
  )
}