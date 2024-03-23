import { RSSItem as RSSItemModel } from "@varsotech/varsoapi/src/app/base";

type RSSItemProps = {
  item: RSSItemModel;
};

function RSSItem({ item }: RSSItemProps) {
  return (
    <div>
      <h4>{item.title}</h4>
    </div>
  );
}

export default RSSItem;
