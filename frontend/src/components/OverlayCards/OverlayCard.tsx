import React from 'react';
import Image from 'next/image';
import styled from 'styled-components';
import { OverlayCardProps } from '@/types/overlayCardProps';

const OverlayCard: React.FC<OverlayCardProps> = ({
    title,
    description,
    onClose,
    imageUrl,
    onLogout,
    logoutButtonLabel,
    children,
}) => (
    <CardContainer>
        <CloseButton onClick={onClose}>
            <Image src="/close-button.svg" alt="Close" width={24} height={24} />
        </CloseButton>
        {imageUrl && (
            <CardImage>
                <Image src={imageUrl} alt={title} width={100} height={100} />
            </CardImage>
        )}
        <CardTitle>{title}</CardTitle>
        <CardDescription>{description}</CardDescription>
        <CardContent>{children}</CardContent>
        {onLogout && logoutButtonLabel && (
            <LogoutButton onClick={onLogout}>{logoutButtonLabel}</LogoutButton>
        )}
    </CardContainer>
);

const CardContainer = styled.article`
    position: relative;
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 100%;
    max-width: 400px;
    max-height: calc(100vh - 40px);
    padding: 24px;
    background-color: #1e1e1e;
    border-radius: 16px;
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.15);
    text-align: center;
    color: #fff;
    transition: all 0.3s ease;
    overflow-y: auto;
    &:hover {
        box-shadow: 0 12px 48px rgba(0, 0, 0, 0.2);
    }
`;

const CardContent = styled.div`
    margin-top: 16px;
    width: 100%;
    overflow-y: auto;
`;

const CloseButton = styled.button`
    position: absolute;
    top: 10px;
    right: 10px;
    background: none;
    border: none;
    cursor: pointer;
    z-index: 1001;
`;

const CardTitle = styled.h2`
    margin-top: 24px;
    font-size: 24px;
    font-weight: 600;
    background: linear-gradient(45deg, #ff6b6b, #feca57);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
`;

const CardDescription = styled.p`
    margin-top: 16px;
    font-size: 16px;
    color: rgba(255, 255, 255, 0.7);
`;

const CardImage = styled.div`
    margin-top: 16px;
`;

const LogoutButton = styled.button`
    margin-top: 24px;
    font-size: 16px;
    font-weight: 600;
    justify-content: center;
    align-items: center;
    border-radius: 12px;
    box-shadow: 0px 1px 2px 0px rgba(0, 0, 0, 0.05);
    background-color: #45b7d8;
    padding: 12px 24px;
    color: #fff;
    cursor: pointer;
    border: none;
    &:hover {
        background-color: #3ca7c8;
    }
`;

export default OverlayCard;
