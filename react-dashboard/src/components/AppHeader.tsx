import React from "react";
import styled from "styled-components";

type AppHeaderProps = {
  title: string;
};

const HeaderContainer = styled.header`
  background-color: #00acd7; /* Golang Blue */
  color: white;
  padding: 1rem;
`;

const Title = styled.h1`
  font-size: 2rem;
  margin: 0;
`;

const AppHeader: React.FC<AppHeaderProps> = ({ title }) => (
  <HeaderContainer>
    <Title>{title}</Title>
  </HeaderContainer>
);

export default AppHeader;
