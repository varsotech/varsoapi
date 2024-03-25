import {
  Organization,
  RSSItem as RSSItemModel,
} from "@varsotech/varsoapi/src/app/base";
import * as Styled from "./RSSItemPreview.style";
import RSSItemMetadata from "./RSSItemMetadata";
import RSSImage from "./RSSItemImage";

type RSSItemPreviewProps = {
  item: RSSItemModel;
  organization: Organization | undefined;
  featured?: boolean;
};

function RSSItemPreview({ item, organization, featured }: RSSItemPreviewProps) {
  return (
    <Styled.RSSItemPreview>
      <div
        style={{
          display: "flex",
          flex: 2,
          flexDirection: "column",
          minWidth: 300,
        }}
      >
        <RSSItemMetadata item={item} organization={organization} />
        <Styled.RSSItemContentContainer href={item.link}>
          <Styled.RSSItemTitle>{item.title}</Styled.RSSItemTitle>
          <Styled.RSSItemDescription>
            {item.description}
          </Styled.RSSItemDescription>
        </Styled.RSSItemContentContainer>
      </div>
      <RSSImage image={item.image} featured={featured} />
    </Styled.RSSItemPreview>
  );
}

export default RSSItemPreview;
