import { useEffect } from "react";
import "./App.css";
import useSWR from "swr";

const fetcher = (url: string): Promise<unknown> =>
  fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      // 'Authorization': 'Bearer your-token' (if needed)
    },
    body: JSON.stringify([
      { suit: "spade", rank: 1 },
      { suit: "spade", rank: 2 },
      { suit: "spade", rank: 3 },
      { suit: "spade", rank: 4 },
      { suit: "spade", rank: 5 },
    ]),
  }).then((res) => res.json());

function App() {
  // const { data, error, isLoading } = useSWR("http://localhost:8080/hand", fetcher);
  useEffect(() => {
    fetch("http://localhost:8080/hand", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify([
        { suit: "spade", rank: 1 },
        { suit: "spade", rank: 2 },
        { suit: "spade", rank: 3 },
        { suit: "spade", rank: 4 },
        { suit: "spade", rank: 5 },
      ]),
    })
      .then((res) => res.json())
      .then(console.log);
  }, []);

  return <div>failed to load</div>;
}

export default App;
