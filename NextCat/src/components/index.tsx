import { useState } from "react";

type Props = {
  initialImageUrl: string;
};

export const MainPage = (props: Props) => {
  const [imageUrl, setImageUrl] = useState(props.initialImageUrl); // 初期値を渡す
  const [loading, setLoading] = useState(false);

  // ボタンをクリックしたときに画像を読み込む処理
  const handleClick = async () => {
    setLoading(true); // 読込中フラグを立てる
    const newImage = await FetchImage();
    setImageUrl(newImage.url); // 画像URLの状態を更新する
    setLoading(false); // 読込中フラグを倒す
  };

  return (
    <div>
      <button onClick={handleClick}>
        他のにゃんこも見る
      </button>
      <div>
        {loading || <img src={imageUrl} />}
      </div>
    </div>
  );
};

type Image = {
  url: string;
};

export const FetchImage = async (): Promise<Image> => {
  const res = await fetch("https://api.thecatapi.com/v1/images/search");
  const images = await res.json();
  console.log(images);
  return images[0];
};
