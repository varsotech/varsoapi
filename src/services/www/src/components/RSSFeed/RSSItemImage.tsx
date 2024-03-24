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
    <div
      style={{
        position: "relative",
      }}
    >
      <img
        src={image?.url}
        alt={image?.title}
        width={featured ? "100%" : 0} // 100% ?
        style={{ flex: 0.5, maxWidth: 700 }}
      />
      <div
        style={{
          position: "absolute",
          top: 0,
          left: 0,
          height: "100%",
          width: "100%",
          backdropFilter: image.blur ? "blur(7px)" : undefined,
        }}
      />
    </div>
  );
}

export default RSSImage;
