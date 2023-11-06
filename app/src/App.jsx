import { createResource } from 'solid-js'

export default function App() {
  const fetchMessage = async () => {
    const res = await fetch(`http://127.0.0.1:3010`)
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