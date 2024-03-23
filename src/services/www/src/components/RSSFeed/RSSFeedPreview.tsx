import {
  RSSFeed as RSSFeedModel,
  RSSItem as RSSItemModel,
} from "@varsotech/varsoapi/src/app/base";
import RSSItemPreview from "./RSSItemPreview";
import * as Styled from "./RSSFeedPreview.style";

type RSSFeedProps = {
  feed: RSSFeedModel;
};

function RSSFeedPreview({ feed }: RSSFeedProps) {
  console.log("feed", feed);
  return (
    <Styled.RSSFeedPreview>
      <Styled.RSSFeedItems>
        {feed?.items?.map((item: RSSItemModel) => {
          return (
            <RSSItemPreview
              key={item.uuid}
              item={item}
              organizationName={feed.title} // TODO: Use organization name from DB
            />
          );
        })}
      </Styled.RSSFeedItems>
    </Styled.RSSFeedPreview>
  );
}

export default RSSFeedPreview;
