
package db

import "os"


type VersionSignature func () (map[string]string, os.Error)


		"username": "phf",
		"password": "somepassword"}
	)

	Note that defaults for all keys are specific to the database
	interface in question and should be documented there.

	The Open() function is free to ignore entries that it has no
	use for. For example, the sqlite3 interface only understands
	"name" and ignores the other well-known keys.

	A database interface is free to introduce additional keys if
	necessary, however those keys have to start with the package
	name of the database interface in question. For example, the
	sqlite3 interface supports the key "sqlite3.vfs".
*/
type OpenSignature func (args map[string]interface{}) (connection Connection, error os.Error)


type Connection interface {

	Prepare(query string) (Statement, os.Error);
	
	Execute(statement Statement, parameters ...) (Cursor, os.Error);
	
	Close() os.Error
}


type InformativeConnection interface {
	Connection;
	
	Changes() (int, os.Error);
}


type FancyConnection interface {
	Connection;
	
	ExecuteDirectly(query string, parameters ...) (*Cursor, os.Error)
}


type TransactionalConnection interface {
	Connection;
	
	Begin() os.Error;
	
	Commit() os.Error;
	
	Rollback() os.Error
}


type Statement interface {
}



type Cursor interface {
	FetchOne() ([]interface {}, os.Error);
	FetchMany(count int) ([][]interface {}, os.Error);
	FetchAll() ([][]interface {}, os.Error);
	Close() os.Error
}

type InformativeCursor interface {
	Cursor;
	Description() (map[string]string, os.Error);
	Results() int;
};

type PythonicCursor interface {
	Cursor;
        FetchDict() (data map[string]interface{}, error os.Error);
        FetchManyDicts(count int) (data []map[string]interface{}, error os.Error);
        FetchAllDicts() (data []map[string]interface{}, error os.Error)
};


