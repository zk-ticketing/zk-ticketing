import React from 'react';
import Image from 'next/image';
import styled from 'styled-components';
import { useRouter } from 'next/router';
import EventCard from '@/components/Events/EventCard';
import { EventListProps } from '@/types/eventListProps';
import withAuth from '@/components/withAuth';

const dummyEvents: EventListProps[] = [
    {
        id: 'event1',
        title: 'Event 1',
        date: '2024-07-08',
        url: 'https://www.example.com',
        description: 'This is a description for event 1.',
    },
    {
        id: 'event2',
        title: 'Event 2',
        date: '2024-07-09',
        url: 'https://www.example.com',
        description: 'This is a description for event 2.',
    },
    {
        id: 'event3',
        title: 'Event 3',
        date: '2024-07-10',
        url: 'https://www.example.com',
        description: 'This is a description for event 3.',
    },
];

const EventsPage: React.FC = () => {
    const router = useRouter();

    const handleRequestTicketCredential = (eventId: string) => {
        console.log('Request ticket credential for event:', eventId);
        // Placeholder logic for requesting a ticket credential
        // Future implementation will include API call to request ticket credential
    };

    const handleScanQRCode = (eventId: string) => {
        console.log('Initiate QR Code Scan for event:', eventId);
        // Navigate to the Scan QR page
        router.push(`/scan-qr?eventId=${eventId}`);
    };

    return (
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
            <EventList>
                {dummyEvents.map((event, index) => (
                    <EventCard
                        key={index}
                        eventId={event.id}
                        eventName={event.title}
                        eventDate={event.date}
                        eventUrl={event.url}
                        eventDescription={event.description}
                        requestTicketCredentialsLabel="Request Credential"
                        onClick={handleRequestTicketCredential}
                        onScanQRCode={handleScanQRCode}
                    />
                ))}
            </EventList>
        </MainContainer>
    );
};

const MainContainer = styled.div`
    background-color: #060708;
    display: flex;
    flex-direction: column;
    max-width: 480px;
    min-height: 100vh;
    margin: 0 auto;
    padding: 15px;
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

const EventList = styled.div`
    display: flex;
    flex-direction: column;
    border-top: 1px solid rgba(255, 255, 255, 0.1);
    margin: 24px 0;
`;

export default withAuth(EventsPage);
