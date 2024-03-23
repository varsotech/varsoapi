import { RSSItem as RSSItemModel } from "@varsotech/varsoapi/src/app/base";
import * as Styled from "./RSSItemPreview.style";
import RSSAuthors from "./RSSAuthors";

type RSSItemPreviewProps = {
  item: RSSItemModel;
  organizationName: string | undefined;
};

function RSSItemPreview({ item, organizationName }: RSSItemPreviewProps) {
  return (
    <Styled.RSSItemPreview>
      <RSSAuthors authors={item.authors} organizationName={organizationName} />
      <a style={{ display: "flex", alignItems: "center" }} href={item.link}>
        <div
          style={{
            paddingRight: 70,
            flex: 1,
          }}
        >
          <a href={item.link}>
            <h3>{item.title}</h3>
            <span
              style={{
                fontSize: 16,
                color: "#5e5e5e",
              }}
            >
              {item.description}
            </span>
          </a>
        </div>
        {item.image?.url ? (
          <img src={item.image?.url} height={100} alt={item.image?.title} />
        ) : null}
      </a>
    </Styled.RSSItemPreview>
  );
}

export default RSSItemPreview;
