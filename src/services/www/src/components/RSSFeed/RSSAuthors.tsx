import {
  Organization,
  RSSAuthor as RSSAuthorModel,
} from "@varsotech/varsoapi/src/app/base";

type RSSAuthorsProps = {
  authors: RSSAuthorModel[];
  organization: Organization | undefined;
};

function RSSAuthors({ authors, organization }: RSSAuthorsProps) {
  return (
    <span>
      {authors
        ?.map((author: RSSAuthorModel) => {
          return author.name;
        })
        .join(", ")}
      {organization ? (
        <span>
          <span style={{ color: "#5c5c5c" }}> from </span>
          <a
            href={organization.websiteUrl}
            style={{
              fontSize: 17,
              fontFamily: "Source Serif Pro",
              fontWeight: 600,
            }}
          >
            {organization.name}
          </a>
        </span>
      ) : (
        ""
      )}
    </span>
  );
}

export default RSSAuthors;
