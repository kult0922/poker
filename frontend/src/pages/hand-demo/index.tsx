import { useEffect, useState } from "react";
import useSWR from "swr";

type Card = {
  suit: string;
  rank: number;
};

type Hand = {
  cards: Card[];
  name: string;
};

const getHandName = (url: string, card: Card[]): Promise<Hand> =>
  fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(card),
  }).then((res) => res.json());

const getFetcher = (url: string): Promise<Card[]> =>
  fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  }).then((res) => res.json());

export const HandDemo = () => {
  const { data, isLoading } = useSWR(
    "http://localhost:8080/community-card",
    getFetcher
  );

  const [hand, setHand] = useState<Hand | undefined>();
  console.log(data);

  useEffect(() => {
    if (data) {
      getHandName("http://localhost:8080/hand", data).then((res) => {
        console.log(res);
        setHand(res);
      });
    }
  }, [data]);

  if (isLoading) return <div>loading...</div>;
  if (data)
    return (
      <div>
        <div>
          <h1 className="">{hand?.name}</h1>
        </div>
        {data.map((e, idx) => {
          return (
            <img
              key={"card-" + idx}
              src={`/cards/${e.rank}_of_${e.suit}s.svg`}
              alt="card"
            />
          );
        })}
      </div>
    );

  return <div>failed to load</div>;
};
