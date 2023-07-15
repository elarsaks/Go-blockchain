import React from "react";
import styled from "styled-components";

const NotificationWrapper = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 0em;
  margin: 1em 0;
  border-radius: 5px;
  color: #333;
  max-width: 817px;
  width: 90%;
  height: 383px;
  overflow: auto;
  background-color: #f2f2f2;
  border: 1px solid #ccc;

  &.info {
    background-color: #00acd7;
    border: 1px solid #066f8a;
  }

  &.warning {
    background-color: #ff9800;
    border: 1px solid #bf7406;
  }

  &.error {
    color: #d94141;
    border: 1px solid #d94141;
  }
`;

const Message = styled.p`
  color: white;
  font-weight: bold;
  margin: 2em;
  font-size: 1.5em;

  &.error {
    color: #d94141;
  }
`;

const SubMessage = styled.h1`
  color: white;
  font-weight: bold;
  margin: 1em;
  font-size: 1.3em;

  &.error {
    color: #d94141;
  }
`;

interface NotificationProps {
  message: string;
  type: "info" | "warning" | "error";
}

const Notification: React.FC<NotificationProps> = ({ message, type }) => {
  if (!message) {
    return null;
  }

  return (
    <NotificationWrapper className={type}>
      <Message className={type}>{message}</Message>
      <SubMessage className={type}>🚧 UNDER DEVELOPMENT 🚧</SubMessage>
    </NotificationWrapper>
  );
};

export default Notification;
