// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

contract Cert {
    address admin;
    event Issued(string indexed course, string id, string grade);

    constructor() {
        admin = msg.sender;
    }

    modifier onlyAdmin() {
        require(msg.sender == admin, "Access Denied");
        _;
    }

    struct Certificate {
        string name;
        string course;
        string grade;
        string date;
    }

    mapping(string => Certificate) public Certificates;

    function issue(
        string memory _id,
        string memory _name,
        string memory _course,
        string memory _grade,
        string memory _date
    ) public onlyAdmin {
        Certificates[_id] = Certificate(_name, _course, _grade, _date);
        emit Issued(_course, _id, _grade);
    }
}
