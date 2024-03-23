import styled from "styled-components";

export const RSSItemPreview = styled.div`
  display: flex;
  /* padding: 25px 20px; */
  /* padding-left: 0; */
  padding-bottom: 25px;
  border-bottom: 1px solid #f2f2f2;
  flex-wrap: wrap-reverse;
  gap: 10px;
  align-items: center;
  justify-content: center;
  /* max-width: 700px; */
`;

export const RSSItemContentContainer = styled.a`
  display: flex;
  flex: 1;
  flex-direction: column;
  /* flex-wrap: wrap; */
`;

export const RSSItemTitle = styled.h3``;

export const RSSItemDescription = styled.span`
  font-size: 16;
  color: #5e5e5e;
`;
