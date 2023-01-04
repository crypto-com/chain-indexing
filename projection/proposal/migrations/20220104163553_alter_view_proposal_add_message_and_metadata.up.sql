ALTER TABLE view_proposals
    ADD msg VARCHAR NULL AFTER data,
    ADD metadata VARCHAR NULL AFTER msg;