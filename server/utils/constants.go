package utils

type InterfaceName int

type InterfaceData struct {
	Value       string
	Description string
	Callable    bool
}

const (
	AddVerified InterfaceName = iota
	Allowance
	Approve
	ApproveAndCall
	AuthorizeOperator
	BalanceOf
	BalanceOf1
	BalanceOfBatch
	Burn
	Burn1
	BurnFrom
	CancelAndReissue
	ImplementInterfaceForAddress
	Cap
	Decimals
	DecreaseAllowance
	DecreaseAllowanceAndCall
	DecreaseSupply
	DefaultOperators
	GetApproved
	GetCurrentFor
	Granularity
	HasHash
	HolderAt
	HolderCount
	IncreaseAllowance
	IncreaseAllowanceAndCall
	IncreaseSupply
	IsApprovedForAll
	IsHolder
	IsOperatorFor
	IsSuperseded
	IsVerified
	Mint
	Name
	OnErc721Received
	OnErc1155BatchReceived
	OnErc1155Received
	OperatorBurn
	OperatorSend
	OwnerOf
	RemoveVerified
	RevokeOperator
	SafeBatchTransferFrom
	SafeTransferFrom
	SafeTransferFrom1
	Send
	SetApprovalForAll
	SupportsInterface
	Symbol
	TokensReceived
	TokensToSend
	TokenByIndex
	TokenOfOwnerByIndex
	TokenUri
	TotalSupply
	Transfer
	Transfer1
	Transfer2
	TransferAndCall
	TransferFrom
	TransferFromAndCall
	UpdateVerified
	URI
)

var InterfaceIdentifiers = map[InterfaceName]InterfaceData{
	AddVerified:                  {Value: "47089f62", Description: "addVerified(address,bytes32)", Callable: false},
	Allowance:                    {Value: "dd62ed3e", Description: "allowance(address,address)", Callable: false},
	Approve:                      {Value: "095ea7b3", Description: "approve(address,uint256)", Callable: false},
	ApproveAndCall:               {Value: "cae9ca51", Description: "approveAndCall(address,uint256,bytes)", Callable: false},
	AuthorizeOperator:            {Value: "959b8c3f", Description: "authorizeOperator(address)", Callable: false},
	BalanceOf:                    {Value: "70a08231", Description: "balanceOf(address)", Callable: false},
	BalanceOf1:                   {Value: "00fdd58e", Description: "balanceOf(address,uint256)", Callable: false},
	BalanceOfBatch:               {Value: "4e1273f4", Description: "balanceOfBatch(address[],uint256[])", Callable: false},
	Burn:                         {Value: "42966c68", Description: "burn(uint256)", Callable: false},
	Burn1:                        {Value: "fe9d9303", Description: "burn(uint256,bytes)", Callable: false},
	BurnFrom:                     {Value: "79cc6790", Description: "burnFrom(address,uint256)", Callable: false},
	CancelAndReissue:             {Value: "79f64720", Description: "cancelAndReissue(address,address)", Callable: false},
	ImplementInterfaceForAddress: {Value: "249cb3fa", Description: "canImplementInterfaceForAddress(bytes32,address)", Callable: false},
	Cap:                          {Value: "355274ea", Description: "cap()", Callable: true},
	Decimals:                     {Value: "313ce567", Description: "decimals()", Callable: true},
	DecreaseAllowance:            {Value: "a457c2d7", Description: "decreaseAllowance(address,uint256)", Callable: false},
	DecreaseAllowanceAndCall:     {Value: "d135ca1d", Description: "decreaseAllowanceAndCall(address,uint256,bytes)", Callable: false},
	DecreaseSupply:               {Value: "869e0e60", Description: "decreaseSupply(uint256,address)", Callable: false},
	DefaultOperators:             {Value: "06e48538", Description: "defaultOperators()", Callable: true},
	GetApproved:                  {Value: "081812fc", Description: "getApproved(uint256)", Callable: false},
	GetCurrentFor:                {Value: "cc397ed3", Description: "getCurrentFor(address)", Callable: false},
	Granularity:                  {Value: "556f0dc7", Description: "granularity()", Callable: true},
	HasHash:                      {Value: "f3221c7f", Description: "hasHash(address,bytes32)", Callable: false},
	HolderAt:                     {Value: "197bc336", Description: "holderAt(uint256)", Callable: false},
	HolderCount:                  {Value: "1aab9a9f", Description: "holderCount()", Callable: true},
	IncreaseAllowance:            {Value: "39509351", Description: "increaseAllowance(address,uint256)", Callable: false},
	IncreaseAllowanceAndCall:     {Value: "5fd42775", Description: "increaseAllowanceAndCall(address,uint256,bytes)", Callable: false},
	IncreaseSupply:               {Value: "124fc7e0", Description: "increaseSupply(uint256,address)", Callable: false},
	IsApprovedForAll:             {Value: "e985e9c5", Description: "isApprovedForAll(address,address)", Callable: false},
	IsHolder:                     {Value: "d4d7b19a", Description: "isHolder(address)", Callable: false},
	IsOperatorFor:                {Value: "d95b6371", Description: "isOperatorFor(address,address)", Callable: false},
	IsSuperseded:                 {Value: "2da7293e", Description: "isSuperseded(address)", Callable: false},
	IsVerified:                   {Value: "b9209e33", Description: "isVerified(address)", Callable: false},
	Mint:                         {Value: "40c10f19", Description: "mint(address,uint256)", Callable: false},
	Name:                         {Value: "06fdde03", Description: "name()", Callable: true},
	OnErc721Received:             {Value: "150b7a02", Description: "onERC721Received(address,address,uint256,bytes)", Callable: false},
	OnErc1155BatchReceived:       {Value: "bc197c81", Description: "onERC1155BatchReceived(address,address,uint256[],uint256[],bytes)", Callable: false},
	OnErc1155Received:            {Value: "f23a6e61", Description: "onERC1155Received(address,address,uint256,uint256,bytes)", Callable: false},
	OperatorBurn:                 {Value: "fc673c4f", Description: "operatorBurn(address,uint256,bytes,bytes)", Callable: false},
	OperatorSend:                 {Value: "62ad1b83", Description: "operatorSend(address,address,uint256,bytes,bytes)", Callable: false},
	OwnerOf:                      {Value: "6352211e", Description: "ownerOf(uint256)", Callable: false},
	RemoveVerified:               {Value: "4487b392", Description: "removeVerified(address)", Callable: false},
	RevokeOperator:               {Value: "fad8b32a", Description: "revokeOperator(address)", Callable: false},
	SafeBatchTransferFrom:        {Value: "2eb2c2d6", Description: "safeBatchTransferFrom(address,address,uint256[],uint256[],bytes)", Callable: false},
	SafeTransferFrom:             {Value: "42842e0e", Description: "safeTransferFrom(address,address,uint256)", Callable: false},
	SafeTransferFrom1:            {Value: "f242432a", Description: "safeTransferFrom(address,address,uint256,uint256,bytes)", Callable: false},
	Send:                         {Value: "9bd9bbc6", Description: "send(address,uint256,bytes)", Callable: false},
	SetApprovalForAll:            {Value: "a22cb465", Description: "setApprovalForAll(address,bool)", Callable: false},
	SupportsInterface:            {Value: "01ffc9a7", Description: "supportsInterface(bytes4)", Callable: false},
	Symbol:                       {Value: "95d89b41", Description: "symbol()", Callable: true},
	TokensReceived:               {Value: "0023de29", Description: "tokensReceived(address,address,address,uint256,bytes,bytes)", Callable: false},
	TokensToSend:                 {Value: "75ab9782", Description: "tokensToSend(address,address,address,uint256,bytes,bytes)", Callable: false},
	TokenByIndex:                 {Value: "4f6ccce7", Description: "tokenByIndex(uint256)", Callable: false},
	TokenOfOwnerByIndex:          {Value: "2f745c59", Description: "tokenOfOwnerByIndex(address,uint256)", Callable: false},
	TokenUri:                     {Value: "c87b56dd", Description: "tokenURI(uint256)", Callable: false},
	TotalSupply:                  {Value: "18160ddd", Description: "totalSupply()", Callable: true},
	Transfer:                     {Value: "a9059cbb", Description: "transfer(address,uint256)", Callable: false},
	Transfer1:                    {Value: "be45fd62", Description: "transfer(address,uint256,bytes)", Callable: false},
	Transfer2:                    {Value: "f6368f8a", Description: "transfer(address,uint256,bytes,string)", Callable: false},
	TransferAndCall:              {Value: "4000aea0", Description: "transferAndCall(address,uint256,bytes)", Callable: false},
	TransferFrom:                 {Value: "23b872dd", Description: "transferFrom(address,address,uint256)", Callable: false},
	TransferFromAndCall:          {Value: "c1d34b89", Description: "transferFromAndCall(address,address,uint256,bytes)", Callable: false},
	UpdateVerified:               {Value: "354b7b1d", Description: "updateVerified(address,bytes32)", Callable: false},
	URI:                          {Value: "0e89341c", Description: "uri(uint256)", Callable: false},
}

type ErcName string
type ErcData []InterfaceName

const (
	Erc20            ErcName = "Erc20"
	Erc20Burnable    ErcName = "Erc20Burnable"
	Erc20Capped      ErcName = "Erc20Capped"
	Erc20Detailed    ErcName = "Erc20Detailed"
	Erc20Mintable    ErcName = "Erc20Mintable"
	Erc20Pausable    ErcName = "Erc20Pausable"
	Erc165           ErcName = "Erc165"
	Erc721           ErcName = "Erc721"
	Erc721Receiver   ErcName = "Erc721Receiver"
	Erc721Metadata   ErcName = "Erc721Metadata"
	Erc721Enumerable ErcName = "Erc721Enumerable"
	Erc820           ErcName = "Erc820"
	Erc1155          ErcName = "Erc1155"
	Erc1155Receiver  ErcName = "Erc1155Receiver"
	Erc1155Metadata  ErcName = "Erc1155Metadata"
	Erc223           ErcName = "Erc223"
	Erc621           ErcName = "Erc621"
	Erc777           ErcName = "Erc777"
	Erc777Receiver   ErcName = "Erc777Receiver"
	Erc777Sender     ErcName = "Erc777Sender"
	Erc827           ErcName = "Erc827"
	Erc884           ErcName = "Erc884"
)

var ErcInterfaceIdentifiers = map[ErcName]ErcData{
	Erc20:            {Allowance, Approve, BalanceOf, TotalSupply, Transfer, TransferFrom},
	Erc20Burnable:    {Burn, BurnFrom},
	Erc20Capped:      {Cap},
	Erc20Detailed:    {Decimals, Name, Symbol},
	Erc20Mintable:    {Mint},
	Erc20Pausable:    {IncreaseAllowance, Approve, DecreaseAllowance, Transfer, TransferFrom},
	Erc165:           {SupportsInterface},
	Erc721:           {Approve, BalanceOf, GetApproved, IsApprovedForAll, OwnerOf, SafeTransferFrom, SafeTransferFrom1, SetApprovalForAll, SupportsInterface, TransferFrom},
	Erc721Receiver:   {OnErc721Received},
	Erc721Metadata:   {Name, Symbol, TokenUri},
	Erc721Enumerable: {TokenByIndex, TokenOfOwnerByIndex, TotalSupply},
	Erc820:           {ImplementInterfaceForAddress},
	Erc1155:          {BalanceOf1, BalanceOfBatch, IsApprovedForAll, SafeBatchTransferFrom, SafeTransferFrom1, SetApprovalForAll},
	Erc1155Receiver:  {OnErc1155BatchReceived, OnErc1155Received},
	Erc1155Metadata:  {URI},
	Erc223:           {BalanceOf, Decimals, Name, Symbol, TotalSupply, Transfer, Transfer1, Transfer2},
	Erc621:           {DecreaseSupply, IncreaseSupply},
	Erc777:           {AuthorizeOperator, BalanceOf, Burn1, DefaultOperators, Granularity, IsOperatorFor, Name, OperatorBurn, OperatorSend, RevokeOperator, Send, Symbol, TotalSupply},
	Erc777Receiver:   {TokensReceived},
	Erc777Sender:     {TokensToSend},
	Erc827:           {ApproveAndCall, DecreaseAllowanceAndCall, IncreaseAllowanceAndCall, TransferAndCall, TransferFromAndCall},
	Erc884:           {AddVerified, CancelAndReissue, GetCurrentFor, HasHash, HolderAt, HolderCount, IsHolder, IsSuperseded, IsVerified, RemoveVerified, Transfer, TransferFrom, UpdateVerified},
}