-- digest(data, type) as pgcrypto defines it, built on PostgreSQL's own
-- sha224/sha256/sha384/sha512/md5. Text input is hashed as its UTF8 bytes,
-- as pgcrypto does in a UTF8 database. Other algorithms raise an error
-- instead of returning a wrong hash.
CREATE FUNCTION digest(data bytea, type text) RETURNS bytea
LANGUAGE plpgsql IMMUTABLE STRICT PARALLEL SAFE AS $$
BEGIN
  CASE lower(type)
    WHEN 'sha224' THEN RETURN sha224(data);
    WHEN 'sha256' THEN RETURN sha256(data);
    WHEN 'sha384' THEN RETURN sha384(data);
    WHEN 'sha512' THEN RETURN sha512(data);
    WHEN 'md5' THEN RETURN decode(md5(data), 'hex');
    ELSE RAISE EXCEPTION 'stutter pgcrypto shim: unsupported digest type %', type;
  END CASE;
END $$;

CREATE FUNCTION digest(data text, type text) RETURNS bytea
LANGUAGE sql IMMUTABLE STRICT PARALLEL SAFE AS $$
  SELECT digest(convert_to(data, 'UTF8'), type)
$$;
