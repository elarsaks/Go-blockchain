import React from "react";
import styled from "styled-components";
import Cube from "./Cube";

type AppHeaderProps = {
  title: string;
};

const HeaderContainer = styled.header`
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: #00acd7;
  color: white;
  padding: 1rem;
`;

const Title = styled.h1`
  font-size: 2rem;
  margin: 0;
  margin-left: 1rem;
`;

const AppHeader: React.FC<AppHeaderProps> = ({ title }) => (
  <HeaderContainer>
    <Title>{title}</Title>
    <Cube />
  </HeaderContainer>
);

export default AppHeader;
