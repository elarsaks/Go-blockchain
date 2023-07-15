import React from "react";
import styled from "styled-components";

const NotificationWrapper = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 1em;
  margin: 0;
  margin-top: 1em;
  width: 90%;
  max-width: 785px;
  border-radius: 5px;
  color: #333;
  overflow: auto;
  background-color: #f2f2f2;
  border: 1px solid #ccc;

  &.info {
    background-color: #00add8;
    border: 1px solid #007d9c;
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
  text-align: center;
  margin: 2em;
  font-size: 1.2em;

  &.error {
    color: #d94141;
  }
`;

const SubMessage = styled.h1`
  color: white;
  font-weight: bold;
  margin: 1em;
  font-size: 1.2em;

  &.error {
    color: black;
  }
`;

interface NotificationProps {
  message: string;
  type: "info" | "warning" | "error";
  underDevelopment?: boolean;
}

const Notification: React.FC<NotificationProps> = ({
  message,
  type,
  underDevelopment,
}) => {
  if (!message) {
    return null;
  }

  return (
    <NotificationWrapper className={type}>
      <Message className={type}>{message}</Message>

      {underDevelopment && (
        <SubMessage className={type}>🚧 UNDER DEVELOPMENT 🚧</SubMessage>
      )}
    </NotificationWrapper>
  );
};

export default Notification;
