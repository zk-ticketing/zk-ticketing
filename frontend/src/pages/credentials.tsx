/* eslint-disable @typescript-eslint/no-explicit-any */

import React, { useState } from 'react';
import styled from 'styled-components';
import { useRouter } from 'next/router';
import Image from 'next/image';
import withAuth from '@/components/withAuth';

const dummyEmailCredential = 'dummy-email-credential';
const dummyTicketCredentials: any[] | (() => any[]) = [];

const CredentialPage: React.FC = () => {
    const router = useRouter();
    const [emailCredential, setEmailCredential] =
        useState(dummyEmailCredential);
    const [ticketCredentials, setTicketCredentials] = useState(
        dummyTicketCredentials,
    );

    const handleRetrieveEmailCredential = () => {
        console.log('Retrieve Email Credential');
        setEmailCredential('new-dummy-email-credential');
    };

    const handleRetrieveTicketCredential = () => {
        console.log('Retrieve Ticket Credential');
        setTicketCredentials([
            ...ticketCredentials,
            { event_id: 'event3', credential: 'dummy-ticket-credential-3' },
        ]);
    };

    return (
        <PageContainer>
            <MainContainer>
                <Header>
                    <GoBackButton onClick={() => router.push('/dashboard')}>
                        <Image
                            src="/left-arrow.svg"
                            alt="go back"
                            width={20}
                            height={20}
                        />
                        <Title>Homepage</Title>
                    </GoBackButton>
                </Header>
                <CredentialSection>
                    <CredentialTitle>Email Credential</CredentialTitle>
                    {emailCredential ? (
                        <CredentialItem>{emailCredential}</CredentialItem>
                    ) : (
                        <NoCredentialText>
                            No email credential found.
                        </NoCredentialText>
                    )}
                    {!emailCredential && (
                        <RetrieveButton onClick={handleRetrieveEmailCredential}>
                            Retrieve Email Credential
                        </RetrieveButton>
                    )}
                </CredentialSection>
                <CredentialSection>
                    <CredentialTitle>Ticket Credentials</CredentialTitle>
                    {ticketCredentials.length > 0 ? (
                        ticketCredentials.map((ticket, index) => (
                            <CredentialItem key={index}>
                                Event ID: {ticket.event_id} - Credential:{' '}
                                {ticket.credential}
                                <GenerateProofButton
                                    onClick={() =>
                                        console.log(
                                            'Generate proof for:',
                                            ticket,
                                        )
                                    }
                                >
                                    Generate Proof
                                </GenerateProofButton>
                            </CredentialItem>
                        ))
                    ) : (
                        <NoCredentialText>
                            No ticket credentials found.
                        </NoCredentialText>
                    )}
                    {!ticketCredentials.length && (
                        <RetrieveButton
                            onClick={handleRetrieveTicketCredential}
                        >
                            Retrieve Ticket Credential
                        </RetrieveButton>
                    )}
                </CredentialSection>
            </MainContainer>
        </PageContainer>
    );
};

const PageContainer = styled.div`
    min-height: 100vh;
    background-color: #060708;
    padding: 24px;
`;

const MainContainer = styled.div`
    color: #fff;
    max-width: 480px;
    margin: 0 auto;
`;

const Header = styled.header`
    display: flex;
    align-items: center;
    background-color: #060708;
    justify-content: center;
`;

const GoBackButton = styled.button`
    background: none;
    border: none;
    cursor: pointer;
    display: flex;
    align-items: center;
    color: #fff;
    font-size: 16px;
    margin-right: 10px;
`;

const Title = styled.span`
    font-size: 18px;
    font-weight: 600;
    background: linear-gradient(45deg, #ff6b6b, #feca57);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
`;

const CredentialSection = styled.div`
    margin: 24px 0;
    padding-top: 24px;
    border-top: 1px solid rgba(255, 255, 255, 0.1);
`;

const CredentialTitle = styled.h2`
    color: rgba(255, 255, 255, 0.8);
    font-size: 20px;
    margin-bottom: 16px;
`;

const CredentialItem = styled.div`
    background-color: rgba(255, 255, 255, 0.1);
    padding: 16px;
    margin: 16px 0;
    border-radius: 12px;
    color: rgba(255, 255, 255, 0.8);
    display: flex;
    justify-content: space-between;
    align-items: center;
    flex-wrap: wrap;
`;

const NoCredentialText = styled.p`
    color: rgba(255, 255, 255, 0.7);
    font-size: 16px;
`;

const Button = styled.button`
    border: none;
    border-radius: 12px;
    padding: 12px 24px;
    cursor: pointer;
    font-weight: 600;
    transition: all 0.3s ease;
    &:hover {
        transform: translateY(-2px);
        box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
    }
    &:active {
        transform: translateY(0);
    }
`;

const RetrieveButton = styled(Button)`
    background: linear-gradient(45deg, #ff6b6b, #feca57);
    color: white;
    margin-top: 16px;
`;

const GenerateProofButton = styled(Button)`
    background: linear-gradient(45deg, #4ecdc4, #45b7d8);
    color: white;
    padding: 8px 16px;
    font-size: 14px;
    margin-top: 8px;
`;

export default withAuth(CredentialPage);
