import { GetServerSideProps, NextPage } from "next";
import { MainPage, FetchImage } from "@/components";

type Props = {
  initialImageUrl: string;
};

const IndexPage: NextPage<Props> = ({ initialImageUrl }) => {
  return (
    <MainPage initialImageUrl={initialImageUrl}/>
  );
};
export default IndexPage;

export const getServerSideProps: GetServerSideProps<Props> = async () => {
  const image = await FetchImage();
  return {
    props: {
      initialImageUrl: image.url,
    },
  };
};
