'use client';

import { Timeline as FbTimeline } from 'flowbite-react';

export const Timeline = () => (
  <FbTimeline>
    <FbTimeline.Item>
      <FbTimeline.Point />
      <FbTimeline.Content>
        <FbTimeline.Time>February 2022</FbTimeline.Time>
        <FbTimeline.Title>Application UI code in Tailwind CSS</FbTimeline.Title>
        <FbTimeline.Body>
          Get access to over 20+ pages including a dashboard layout, charts, kanban board, calendar, and pre-order
          E-commerce & Marketing pages.
        </FbTimeline.Body>
      </FbTimeline.Content>
    </FbTimeline.Item>
    <FbTimeline.Item>
      <FbTimeline.Point />
      <FbTimeline.Content>
        <FbTimeline.Time>March 2022</FbTimeline.Time>
        <FbTimeline.Title>Marketing UI design in Figma</FbTimeline.Title>
        <FbTimeline.Body>
          All of the pages and components are first designed in Figma and we keep a parity between the two versions even
          as we update the project.
        </FbTimeline.Body>
      </FbTimeline.Content>
    </FbTimeline.Item>
  </FbTimeline>
);