import React from "react";
import styled from "styled-components";

const FooterContainer = styled.footer`
  background-color: #00acd7;
  color: white;
  padding: 1rem;
  display: flex;
  justify-content: center;
  align-items: center;
`;

const Link = styled.a`
  margin: 0 1rem;
  color: white; /* Make the link text white */
  text-decoration: none;
`;

type FooterProps = {
  githubUrl: string;
  linkedinUrl: string;
  websiteUrl: string;
};

const Footer: React.FC<FooterProps> = ({
  githubUrl,
  linkedinUrl,
  websiteUrl,
}) => (
  <FooterContainer>
    <Link href={githubUrl} target="_blank" rel="noopener noreferrer">
      GitHub
    </Link>
    <Link href={linkedinUrl} target="_blank" rel="noopener noreferrer">
      LinkedIn
    </Link>
    <Link href={websiteUrl} target="_blank" rel="noopener noreferrer">
      Website
    </Link>
  </FooterContainer>
);

export default Footer;
