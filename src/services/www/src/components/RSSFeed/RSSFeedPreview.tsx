import {
  Organization,
  RSSFeed as RSSFeedModel,
  RSSItem as RSSItemModel,
} from "@varsotech/varsoapi/src/app/base";
import RSSItemPreview from "./RSSItemPreview";
import * as Styled from "./RSSFeedPreview.style";

type RSSFeedProps = {
  feed: RSSFeedModel | undefined;
  organization: Organization | undefined;
  featured?: boolean;
};

function RSSFeedPreview({ feed, organization, featured }: RSSFeedProps) {
  return (
    <Styled.RSSFeedPreview>
      <Styled.RSSFeedItems>
        {feed?.items?.map((item: RSSItemModel, i: number) => {
          if (
            featured &&
            item.title !==
              "Malcolm X’s final written words were about Zionism. Here is what he said."
          )
            return;

          return (
            <RSSItemPreview
              key={item.uuid}
              item={item}
              organization={organization}
              featured={
                featured ===
                (item.title ===
                  "Malcolm X’s final written words were about Zionism. Here is what he said.")
              }
            />
          );
        })}
      </Styled.RSSFeedItems>
    </Styled.RSSFeedPreview>
  );
}

export default RSSFeedPreview;
