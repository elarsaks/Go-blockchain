import React from "react";
import styled from "styled-components";

type ButtonProps = {
  $primary?: boolean;
};

export const Button = styled.button<ButtonProps>`
  /* Adapt the colors based on primary prop */
  background: ${(props) => (props.$primary ? "#00acd7" : "white")};
  color: ${(props) => (props.$primary ? "white" : "#00acd7")};

  font-size: 1em;
  margin: 1em;
  padding: 0.25em 1em;
  border: 2px solid #00acd7; /* Golang Blue */
  border-radius: 3px;
`;

export const ParentComponent: React.FC = () => (
  <div>
    <Button>Normal</Button>
    <Button $primary>Primary</Button>
  </div>
);
