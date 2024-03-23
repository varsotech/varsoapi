import { RSSImage as RSSImageModel } from "@varsotech/varsoapi/src/app/base";

type RSSImageProps = {
  image: RSSImageModel | undefined;
  featured?: boolean;
};

function RSSImage({ image, featured }: RSSImageProps) {
  if (!image?.url) {
    return null;
  }
  return (
    <img
      src={image?.url}
      alt={image?.title}
      // height={featured ? "100%" : 200}
      width={featured ? "100%" : 0} // 100% ?
      style={{ flex: 1, maxWidth: 700 }}
    />
  );
}

export default RSSImage;
