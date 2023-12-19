import React from 'react';
import { render } from "@testing-library/react";
import { MainPage } from "../components/index";

const initialImageUrl = "https://cdn2.thecatapi.com/images/7qr.jpg";

describe("コンポーネントのテスト", () => {
  it("コンポーネントのレンダリングのテスト", async () => {
    render(<MainPage
      initialImageUrl={initialImageUrl}
    />);
  });
});
