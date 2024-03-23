import { RSSAuthor as RSSAuthorModel } from "@varsotech/varsoapi/src/app/base";

type RSSAuthorsProps = {
  authors: RSSAuthorModel[];
  organizationName: string | undefined;
};

function RSSAuthors({ authors, organizationName }: RSSAuthorsProps) {
  return (
    <div style={{ marginBottom: 8, display: "inline-block" }}>
      {authors
        .map((author: RSSAuthorModel) => {
          return author.name;
        })
        .join(", ")}
      {organizationName ? (
        <span>
          <span style={{ color: "#5c5c5c" }}> from </span>
          {organizationName || ""}
        </span>
      ) : (
        ""
      )}
    </div>
  );
}

export default RSSAuthors;
