import * as Styled from "./Layout.style";

type LayoutProps = {
  children: React.ReactNode;
};

function Layout({ children }: LayoutProps) {
  return (
    <Styled.Layout>
      <Styled.Navbar />
      <Styled.PageWrapper>{children}</Styled.PageWrapper>
    </Styled.Layout>
  );
}

export default Layout;
