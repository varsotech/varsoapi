import { RSSFeed as RSSFeedModel } from "@varsotech/varsoapi/src/app/base";
import RSSFeed from "../../components/RSSFeed/RSSFeed";
import { useNews } from "../../api/news";

function Home() {
  const { data } = useNews();

  return (
    <div>
      <h1>News</h1>
      {data?.data?.feeds.map((feed: RSSFeedModel) => {
        return <RSSFeed key={feed.title} feed={feed} />;
      }) || null}
    </div>
  );
}

export default Home;
